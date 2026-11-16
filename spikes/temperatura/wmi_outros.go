//go:build !windows

package main

func testarWMI() bool {
	avisar("Não se aplica: este caminho só existe no Windows.\n")
	return false
}
