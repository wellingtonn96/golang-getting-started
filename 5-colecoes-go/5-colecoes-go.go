package main

import (
	"fmt"
	"net/http"
	"os"
	"reflect"
	"time"
)

const monitoramento = 5
const delay = 6

func main() {
	exibeIntroducao()

	// exibeNomes()
	
	for {
		exibeMenu()

		// _, idade := devolveNomeEIdade()

		// fmt.Println(nome)
		// fmt.Println(idade)

		comando := leComando()

		switch comando {
			case 1:
				iniciarMonitoramento()
			case 2:
				fmt.Println("Exibindo Logs...")
			case 3:
				fmt.Println("Saindo do programa")
			default:
				fmt.Print("Não conheço este comando")
				os.Exit(-1)
		}
	}
}

func exibeIntroducao() {
	nome := "Douglas"
	idade := 24
	versao := 1.1
	fmt.Println("Olá, sr.", nome)
	fmt.Println("Olá , sr. ", nome, "Sua idade idade é ", idade)
	fmt.Println("Este programa está na versão", versao)
	fmt.Println("O tipo da variavel nome é", reflect.TypeOf(nome))
	fmt.Println("O tipo da variavel idade", reflect.TypeOf(versao))
}

func leComando() int {
	var comandoLido int
	fmt.Scanf("%d", &comandoLido)
	fmt.Println("O endereço desta variavel comando é: ", &comandoLido)

	return comandoLido
}

func exibeMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do programa")
}

func devolveNomeEIdade() (string, int) {
	nome := "wellington"
	idade := 24
	return nome, idade
}	

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	sites := []string{ "https://httpbin.org/status/404", "https://www.alura.com.br", "https://www.caelum.com.br" }

	fmt.Println((sites))
	for i := 0; i < monitoramento; i ++ {
		for i, site := range sites {
			fmt.Println("Testando site", i, ":", site)
			testeSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}
	fmt.Println("")
}

func testeSite(site string) {
	response, _ := http.Get(site)

	if response.StatusCode == 200 {
		fmt.Println("Site", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, "esta com problemas, satus code", response.StatusCode)
	}
}

// func exibeNomes() {
// 	nomes := []string{"Douglas", "Daniel", "Bernardo"}
// 	fmt.Println(nomes)
// 	fmt.Println(reflect.TypeOf(nomes))
// 	fmt.Println("O meu slice tem", len(nomes), "items")
// 	fmt.Println("O meu slice tem a capacidade para", cap(nomes), "items")
	
// 	nomes = append(nomes, "Aparecida")

// 	fmt.Println(nomes)
// 	fmt.Println(reflect.TypeOf(nomes))
// 	fmt.Println("O meu slice tem", len(nomes))
// 	fmt.Println("O meu slice tem a capacidade para", cap(nomes), "items")
// }
