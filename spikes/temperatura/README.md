# Spike: de onde vem a temperatura do processador?

**Pergunta:** no Windows, de onde o agente (em Go) consegue tirar a temperatura do processador?

É o maior risco técnico do MVP: o agente inteiro depende dessa leitura. O Windows não tem
um jeito padrão e confiável de fornecer esse valor, então o teste compara dois caminhos:

1. **Leitura direta pelo Windows (WMI).** Não precisa instalar nada, mas costuma exigir
   administrador e muitas vezes devolve a temperatura da placa-mãe, não do processador.
2. **LibreHardwareMonitor.** Programa gratuito que lê os sensores de verdade e pode
   disponibilizar os valores em `http://localhost:8085/data.json`.

## Como rodar

Pré-requisito: [Go](https://go.dev/dl/) instalado.

```bash
cd spikes/temperatura
go mod tidy
go run .
```

Rode duas vezes: uma normal e outra num terminal **como administrador**.

Para o caminho 2: baixe o LibreHardwareMonitor no GitHub oficial do projeto, abra como
administrador e ligue `Options > Remote Web Server > Run`. Depois rode o teste de novo.

## Resultado

> Preencher depois de rodar.

- Máquina / processador:
- Caminho 1 (WMI) sem administrador:
- Caminho 1 (WMI) como administrador:
- Caminho 2 (LibreHardwareMonitor):
- Os valores batem com outro programa?
- Decisão (vira a ADR 0006):
