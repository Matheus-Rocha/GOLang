package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const delay = 6
const monitoring_times = 5

func main() {
	fmt.Println("Starting the algorithm...")
	fmt.Println("")

	scanner := bufio.NewScanner(os.Stdin)
	for {
		LoadMenu()
		scanner.Scan()
		option := scanner.Text()
		switch option {
		case "1":
			list_sites := []string{
				"https://github.com/",
				"https://www.mercadobitcoin.com.br/",
				"https://random-status-code.herokuapp.com/"}

			log_option := LogOption()
			StartMonitoring(list_sites, log_option)
			break
		case "2":
			var sites_list []string
			scanner_qtd_sites := bufio.NewScanner(os.Stdin)
			fmt.Println("How many sites: ")
			scanner_qtd_sites.Scan()
			string_qtd_sites := scanner_qtd_sites.Text()

			qtd_sites, err := strconv.Atoi(string_qtd_sites)
			if err != nil {
				log.Println("Error: ", err)
			}
			site_scanner := bufio.NewScanner(os.Stdin)
			for i := 0; i < qtd_sites; i++ {
				fmt.Println("site url: ")
				site_scanner.Scan()
				site := site_scanner.Text()
				sites_list = append(sites_list, site)
			}

			log_option := LogOption()
			StartMonitoring(sites_list, log_option)
			break
		case "3":
			list_sites := ReadFile()
			log_option := LogOption()
			StartMonitoring(list_sites, log_option)
			break
		case "4":
			ShowLogs()
			break
		case "9":
			os.Exit(-1)
		default:
			fmt.Println("Entrada inválida")
		}
	}
}

func testing(site string, make_log bool) {
	response, err := http.Get(site)
	if err != nil {
		log.Println("Error: ", err)
	}
	if response.StatusCode == 200 && make_log == false {
		fmt.Println("Site sucessfully loaded!")
	}
	if response.StatusCode == 200 && make_log == true {
		fmt.Println("Site sucessfully loaded!")
		MakeLog(site, make_log)
	}
	if response.StatusCode != 200 && make_log == true {
		fmt.Println("Site load failed: StatusCode: ", response.StatusCode)
		MakeLog(site, make_log)
	}
	if response.StatusCode != 200 && make_log == false{
		fmt.Println("Site load failed: StatusCode: ", response.StatusCode)
	}
}

func StartMonitoring(sites []string, has_logger bool) {
	fmt.Println("Starting monitoring...")

	for i := 0; i < monitoring_times; i++ {
		for j, site := range sites {
			fmt.Println(j+1, "- site: ", site)
			testing(site, has_logger)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}
	fmt.Println("Monitoring finished!")
}

func ReadFile() []string {
	fmt.Println("Reading file...")

	var sites []string

	file, err := os.Open("sites_file.txt")
	if err != nil {
		log.Println("Error: ", err)
	}

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		// TrimSpace is used to clear specials chars like \n
		line = strings.TrimSpace(line)
		sites = append(sites, line)

		if err == io.EOF {
			break
		}
	}
	return sites
}

func LoadMenu() {
	fmt.Println("1 - Monitoring using default sites on algorithm")
	fmt.Println("2 - Monitoring using personalized sites")
	fmt.Println("3 - Monitoring reading a sites_file.txt")
	fmt.Println("4 - Show logs")
	fmt.Println("9 - Leave")
}

func LogOption() bool {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Do u want to create log file and save data?\n1 - Sim\n2 - Não")
	scanner.Scan()
	log_option := scanner.Text()
	if log_option == "1" {
		return true
	} else {
		return false
	}
}

func MakeLog(site string, status bool) {

	file, err := os.OpenFile("logs.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		log.Println("Error: ", err)
	}

	file.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " status: " + strconv.FormatBool(status) + "\n")

	file.Close()
}

func ShowLogs() {
	file, err := os.Open("logs.txt")
	if err != nil {
		log.Println("Error: ", err)
	}

	reader := bufio.NewReader(file)
	for {
		logs_list, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}

		fmt.Println(logs_list)
	}
	file.Close()
}
