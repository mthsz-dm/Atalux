// Teste descartável: de onde o agente consegue tirar a temperatura do processador?
//
// Caminho 1: leitura direta pelo Windows (WMI, classe MSAcpi_ThermalZoneTemperature).
// Caminho 2: LibreHardwareMonitor com o "Remote Web Server" ligado (http://localhost:8085).
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"runtime"
	"strings"
	"time"
)

const lhmURL = "http://localhost:8085/data.json"

func main() {
	fmt.Printf("Sistema: %s/%s\n\n", runtime.GOOS, runtime.GOARCH)

	fmt.Println("== Caminho 1: leitura direta pelo Windows (WMI) ==")
	ok1 := testarWMI()

	fmt.Println("\n== Caminho 2: LibreHardwareMonitor ==")
	ok2 := testarLHM()

	fmt.Println("\n== Resultado ==")
	switch {
	case ok2:
		fmt.Println("LibreHardwareMonitor funcionou. É o caminho mais confiável: costuma trazer a temperatura real de cada núcleo.")
	case ok1:
		fmt.Println("Só a leitura direta funcionou. Compare o valor com outro programa: no Windows ela às vezes não é a temperatura real do processador.")
	default:
		fmt.Println("Nenhum caminho funcionou. Tente rodar como administrador e confira se o LibreHardwareMonitor está aberto com o servidor web ligado.")
	}
	fmt.Println("Anote o resultado em spikes/temperatura/README.md.")
}

// ---------- Caminho 2: LibreHardwareMonitor ----------

// No arquivo data.json, cada item da árvore de hardware é um "nó" com texto, valor e filhos.
type noLHM struct {
	Text     string  `json:"Text"`
	Value    string  `json:"Value"`
	Children []noLHM `json:"Children"`
}

type sensor struct {
	caminho string
	valor   string
}

func testarLHM() bool {
	cliente := http.Client{Timeout: 3 * time.Second}
	resp, err := cliente.Get(lhmURL)
	if err != nil {
		fmt.Println("Não conectou em", lhmURL)
		fmt.Println("Abra o LibreHardwareMonitor e ligue: Options > Remote Web Server > Run.")
		return false
	}
	defer resp.Body.Close()

	var raiz noLHM
	if err := json.NewDecoder(resp.Body).Decode(&raiz); err != nil {
		fmt.Println("Conectou, mas não entendeu a resposta:", err)
		return false
	}

	var temps []sensor
	coletar(raiz, nil, &temps)
	if len(temps) == 0 {
		fmt.Println("Conectou, mas não encontrou nenhum sensor de temperatura.")
		return false
	}
	fmt.Printf("Encontrou %d sensores de temperatura:\n", len(temps))
	for _, s := range temps {
		marca := "  "
		if pareceProcessador(s.caminho) {
			marca = "→ " // provavelmente do processador
		}
		fmt.Printf("%s%-60s %s\n", marca, s.caminho, s.valor)
	}
	fmt.Println("(→ = provavelmente do processador)")
	return true
}

func coletar(n noLHM, caminho []string, saida *[]sensor) {
	atual := caminho
	if n.Text != "" {
		atual = append(append([]string{}, caminho...), n.Text)
	}
	if strings.HasSuffix(strings.TrimSpace(n.Value), "°C") && len(n.Children) == 0 {
		*saida = append(*saida, sensor{strings.Join(atual, " > "), n.Value})
	}
	for _, f := range n.Children {
		coletar(f, atual, saida)
	}
}

func pareceProcessador(caminho string) bool {
	c := strings.ToLower(caminho)
	for _, p := range []string{"cpu", "intel", "amd ryzen", "core #", "package", "tctl"} {
		if strings.Contains(c, p) && !strings.Contains(c, "gpu") {
			return true
		}
	}
	return false
}

// Usado pelas versões de WMI.
func avisar(format string, a ...any) { fmt.Fprintf(os.Stdout, format, a...) }
