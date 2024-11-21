package main

import (
	"fmt"
	"reflect"
)

func main() {
	nome := "Douglas"
	idade := 24
	versao := 1.1
	fmt.Println("Olá, sr.", nome)
	fmt.Println("Olá , sr. ", nome, "Sua idade idade é ", idade)
	fmt.Println("Este programa está na versão", versao)
	fmt.Println("O tipo da variavel nome é", reflect.TypeOf(nome))
	fmt.Println("O tipo da variavel idade", reflect.TypeOf(versao))

	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do programa")

	var comando int
	fmt.Scanf("%d", &comando)
	fmt.Println("O endereço desta variavel comando é: ", &comando)
	fmt.Println("O comando escolhido foi", comando)

	if comando == 1 {
		fmt.Println("Monitorando...")
	} else if comando == 2 {	
		fmt.Println("Exibindo Logs...")
	} else if comando == 0 {
		fmt.Println("Saindo do programa")
	} else {
		fmt.Println("Não conheço este comando")
	}

	switch comando {
		case 1:
			fmt.Println("Monitorando")
		case 2:
			fmt.Println("Exibindo Logs...")
		case 3:
			fmt.Println("Saindo do programa")
		default:
			fmt.Print("Não conheço este comando")
	}
}