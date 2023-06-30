package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	exibeIntro()
	for {
		entrada := input()
		switch entrada {
		case 1:
			fmt.Println("Monitorando....")
			monitoramento()
		case 2:
			fmt.Println("Exibindo logs....")
		case 0:
			fmt.Println("Saindo....")
			os.Exit(0)
		default:
			fmt.Println("Não conheço este comando...")
			os.Exit(1)
		}
	}
}

func exibeIntro() {
	nome := "vinicius"
	versao := 1.1

	fmt.Println("hello world")
	fmt.Println("olá sr.", nome)
	fmt.Println("este programa esta na versao", versao)
}

func input() int {
	fmt.Println("1 - iniciar monitoramento")
	fmt.Println("2 - exibir os logs")
	fmt.Println("0 - sair")

	var cmd int
	fmt.Scan(&cmd)

	fmt.Printf("input recebido: %d \n", cmd)
	return cmd
}

func monitoramento() {
	sites := readTxt()
	for {
		for i, site := range sites {
			fmt.Print("[*] Testando site ", i, ": ", site)
			status, _ := testaSite(site)
			log(site, status)
		}
		fmt.Println("-----------------------------------------------------")
		time.Sleep(1 * time.Minute)
	}
}

func testaSite(site string) (bool, int) {
	resp, _ := http.Get(site)
	if resp.StatusCode == 200 {
		fmt.Println(" [+] (status 200)")
		return true, resp.StatusCode
	} else {
		fmt.Println("[-] ", resp.Status)
		return false, resp.StatusCode
	}
}

func readTxt() []string {
	res, _ := os.Open("sites.txt")
	leitor := bufio.NewReader(res)
	var sites []string
	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		sites = append(sites, linha)
		if err == io.EOF {
			break
		}
	}
	return sites
}

func log(site string, status bool) {
	arquivo, erro := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if erro != nil {
		fmt.Println(erro)
	}
	arquivo.WriteString(site + " - " + "online: " + strconv.FormatBool(status) + "\n")
	arquivo.Close()
}
