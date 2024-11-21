package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

const monitoramento = 5
const delay = 6

func main() {
	exibeIntroducao()
	registraLog("site-false", false)

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
				imprimeLogs()
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
	// sites := []string{ "https://httpbin.org/status/404", "https://www.alura.com.br", "https://www.caelum.com.br" }

	sites := leSitesDoArquivo()

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
	response, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro: ", err)
	}

	if response.StatusCode == 200 {
		fmt.Println("Site", site, "foi carregado com sucesso!")
		registraLog(site, true)
	} else {
		fmt.Println("Site:", site, "esta com problemas, satus code", response.StatusCode)
		registraLog(site, false)
	}
}

func leSitesDoArquivo() []string {
	var sites []string
	arquivo, err := os.Open("7-escrevendo-arquivos/sites.txt")

	// arquivo, err := ioutil.ReadFile("7-escrevendo-arquivos/sites.txt")
	
	if err != nil {
		fmt.Println("Ocorreu um erro: ", err)
	}

	leitor := bufio.NewReader(arquivo)

	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)

		sites = append(sites, linha)

		if err == io.EOF {
			break
		}
	}

	arquivo.Close()
	
	fmt.Println(sites)
	return sites
}

func registraLog(site string, status bool) {
	arquivo, err := os.OpenFile("7-escrevendo-arquivos/log.txt", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(arquivo)

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + "- online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}

func imprimeLogs() {
	arquivo, err := ioutil.ReadFile("7-escrevendo-arquivos/log.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(arquivo))
}