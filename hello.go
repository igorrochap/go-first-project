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

const monitoringTime = 5
const delay = 5

func main() {
	showHelloMessage()

	for {
		showMenu()

		command := scanCommand()

		switch command {
		case 1:
			startMonitoring()
		case 2:
			showLogs()
			fmt.Println("Showing logs...")
		case 0:
			fmt.Println("Exiting...")
			os.Exit(0)
		default:
			fmt.Println("Invalid instruction!")
			os.Exit(-1)
		}
	}
}

func showHelloMessage() {
	name := "Igor"
	version := 1.1 // float32
	fmt.Println("Hello Mr.", name)
	fmt.Println("Version:", version)
}

func showMenu() {
	fmt.Println("\n1 - Start monitoring")
	fmt.Println("2 - Show logs")
	fmt.Println("0 - Exit")
}

func scanCommand() int {
	var command int
	fmt.Scanf("%d", &command)
	return command
}

func startMonitoring() {
	fmt.Println("\nStarting monitoring...")
	// sites := []string{"https://www.alura.com.br", "https://random-status-code.herokuapp.com", "https://www.google.com"}
	sites := readFileSites()

	for i := 0; i < monitoringTime; i++ {
		fmt.Println("BATCH", i+1)

		for _, site := range sites {
			testSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}

	// for i := 0; i < len(sites); i++ {
	// code here
	// }
}

func readFileSites() []string {
	var sites []string
	file, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("An error has ocurred: ", err)
	}

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		sites = append(sites, strings.TrimSpace(line))
		if err == io.EOF {
			break
		}
	}

	file.Close()

	return sites
}

func testSite(site string) {
	response, err := http.Get(site)
	if err != nil {
		fmt.Println("An error has ocurred: ", err)
	}

	if response.StatusCode == 200 {
		fmt.Println("Site:", site, "successfully loaded!")
		writeLog(site, true)
	} else {
		fmt.Println("Site:", site, "is having problems! CODE:", response.StatusCode)
		writeLog(site, false)
	}
}

func writeLog(site string, status bool) {
	logFile, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("An error has ocurred: ", err)
	}

	now := time.Now().Format("02/01/2006 15:04:05")
	logFile.WriteString(now + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")

	logFile.Close()
}

func showLogs() {
	file, err := ioutil.ReadFile("log.txt")
	if err != nil {
		fmt.Println("An error has ocurred: ", err)
	}
	fmt.Println(string(file))
}
