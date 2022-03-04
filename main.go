package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
)

var _token string
var _listenPort string
var data []symbolData
var symbols []string

type symbolData struct {
	dayChange string
	name      string
	open      float64
	prevClose float64
	price     float64
	symbol    string
	time      int64
	todayHigh float64
	todayLow  float64
}

func getEnvironmentVariables() {
	fileData, err := ioutil.ReadFile("./env")
	if err != nil {
		fmt.Println("env file not found")
		return
	}
	lines := strings.Split(string(fileData), "\n")
	for _, line := range lines {
		if len(line) == 0 {
			break
		}
		spl := strings.Split(line, "=")
		key, val := spl[0], spl[1]
		os.Setenv(key, val)
	}
}

func requestBuilder(symbol string) string {
	sited := "http://198.58.126.207:5000/raw/%s?token=%s"
	token := "oeiprjqoeqk24k9"
	return fmt.Sprintf(sited, symbol, token)
}

func requestData(url string) string {
	resp, _ := http.Get(requestBuilder("NFLX"))
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}

func complete() {
	// for _, sym := range symbols {
	// 	//rd := requestData(requestBuilder(sym))

	// }
}

func handleRoutes() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/complete", complete)
	log.Fatal(http.ListenAndServe(":"+_listenPort, router))
}

func main() {
	getEnvironmentVariables()
	_token = os.Getenv("TOKEN")
	_listenPort = os.Getenv("LISTENPORT")
	handleRoutes()
}
