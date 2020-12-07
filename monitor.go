package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const loopingMonitor = 5
const delay = 10

func main() {

	helloMessage()
	// for without condition is equal to while(true)
	for {
		showMenu()
		command := readInput()

		switch command {
		case 1:
			initMonitor()
		case 2:
			fmt.Println("load logs ...")
			showLogs()
		case 0:
			fmt.Println("Bye :)")
			os.Exit(0)
		default:
			fmt.Println("invalid command")
			os.Exit(-1)
		}
	}
}

func helloMessage() {
	//string e int podem ser inferidos
	// var name = "string" é igual a name:= "string" (obs: usar essa redução implica em usar a inferência)
	name := "Douglas"

	// float pode ser inferido também, mas não é recomendado por existirem 2 tipos : 32 e 64
	version := 1.1

	// o "," concatena no print
	fmt.Println("Hello", name)
	fmt.Println("Program Version", version)
}

func readInput() int {
	var readCommand int
	//fmt.Scanf("%d", &command) %d diz que só aceita inteiro || & indica o endereço de memória da variável que o acompanha
	fmt.Scan(&readCommand) //fmt.Scan (sem f) entende o tipo de variável aceito por inferência
	return readCommand
}

func showMenu() {
	fmt.Println("1 - Start monitoring")
	fmt.Println("2 - Show logs")
	fmt.Println("0 - Exit")
}

func initMonitor() {
	fmt.Println("monitoring...")
	sites := getTargets()

	// traditional FOR
	// for i := 0; i < len(sites); i++ {
	// 	fmt.Println(sites[i])
	// }

	//Range returns index and element for each iteration
	for i := 0; i < loopingMonitor; i++ {
		for index, site := range sites {
			fmt.Println("testing site ", index+1, ": ", site)
			testSite(site)
			fmt.Println()

		}
		time.Sleep(delay * time.Second)
	}
}

func testSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("error occurred: ", err)
	}
	if resp.StatusCode == 200 {
		fmt.Println("site ", site, "successfully access ")
		writeLog(site, true)
	} else {
		fmt.Println("site ", site, " it has problems. Status Code: ", resp.StatusCode)
		writeLog(site, false)
	}
}

func getTargets() []string {
	var sites []string
	file, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("error occurred: ", err)
	}

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		sites = append(sites, line)
		if err == io.EOF {
			break
		}
	}
	file.Close()
	return sites
}

func writeLog(site string, status bool) {

	file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(file)
	}

	file.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")
	file.Close()

}

func showLogs() {
	file, err := ioutil.ReadFile("log.txt")
	fmt.Println(string(file))
	if err != nil {
		fmt.Println(err)
	}

}
