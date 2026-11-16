//go:build windows

package main

import "github.com/yusufpapurcu/wmi"

// Classe do Windows com as "zonas térmicas" da placa-mãe.
// CurrentTemperature vem em décimos de Kelvin.
type MSAcpi_ThermalZoneTemperature struct {
	InstanceName       string
	CurrentTemperature uint32
}

func testarWMI() bool {
	var zonas []MSAcpi_ThermalZoneTemperature
	err := wmi.QueryNamespace("SELECT InstanceName, CurrentTemperature FROM MSAcpi_ThermalZoneTemperature", &zonas, `root\WMI`)
	if err != nil {
		avisar("Falhou: %v\n", err)
		avisar("Costuma exigir administrador, e muitos computadores simplesmente não oferecem essa informação.\n")
		return false
	}
	if len(zonas) == 0 {
		avisar("O Windows respondeu, mas sem nenhuma zona térmica.\n")
		return false
	}
	for _, z := range zonas {
		celsius := float64(z.CurrentTemperature)/10 - 273.15
		avisar("%s: %.1f °C\n", z.InstanceName, celsius)
	}
	avisar("Atenção: essas zonas são da placa-mãe e nem sempre são a temperatura real do processador.\n")
	return true
}
