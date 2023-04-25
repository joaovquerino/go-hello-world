package main

import (
	"fmt"
	"os"
	//"reflect"
)

// toda funcao que vem de um pacote externo começa com letra maiuscula em GO --> nomepacote.Funcao()
func main() {

	exibeIntro()
	exibeMenu()
	comando := leComando()

	// if comando == 1 {
	// 	fmt.Println("Monitorando [...]")
	// } else if comando == 2 {
	// 	fmt.Println("Exibindo log [...]")
	// } else if comando == 0 {
	// 	fmt.Println("Encerrando programa.")
	// } else {
	// 	fmt.Println("Comando não reconhecido.")
	// }

	switch comando {
	case 1:
		fmt.Println("Monitorando [...]")
	case 2:
		fmt.Println("Exibindo log [...]")
	case 0:
		fmt.Println("Encerrando programa.")
		os.Exit(0)
	default:
		fmt.Println("Comando não reconhecido.")
	}
}

func exibeIntro() {
	nome := "João"
	idade := 25
	versao := 1.1
	fmt.Println("Olá, Sr.", nome, "sua idade é", idade)
	fmt.Println("Este programa está na versão", versao)
}

func exibeMenu() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")
	fmt.Println("--- Digite a opção escolhida e pressione ENTER ---")
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	// fmt.Println(reflect.TypeOf(comando))
	fmt.Println("O comando escolhido foi", comandoLido)

	return comandoLido
}
