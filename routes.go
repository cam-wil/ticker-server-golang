package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func HandleRoutes() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/complete", Complete)
	log.Fatal(http.ListenAndServe(":"+_listenPort, router))

}

func Timed() {
	var tempData SymbolData
	var tempArr []SymbolData

	for sym, _ := range symbols {
		rd := RequestData(RequestBuilder(sym))
		err := json.Unmarshal(rd, &tempData)
		if err != nil {
			log.Fatal("failed to unmarshal", string(rd))
		}
		fmt.Println(sym, "got")
		tempArr = append(tempArr, tempData)
		time.Sleep(1 * time.Second)
	}
	SymbolStore = tempArr
	thisPrice := calculatePrice(SymbolStore)
	TempComplete = CompleteData{Worth: thisPrice, Errors: setErrors(), MaxWorth: setMax(thisPrice), MinWorth: setMin(thisPrice), Stocks: SymbolStore}
}

// /complete returns worth and all ticker information
func Complete(w http.ResponseWriter, r *http.Request) {
	// var tempData SymbolData
	// var tempArr []SymbolData

	// for sym, _ := range symbols {
	// 	rd := RequestData(RequestBuilder(sym))
	// 	err := json.Unmarshal(rd, &tempData)
	// 	if err != nil {
	// 		log.Fatal("failed to unmarshal", string(rd))
	// 	}
	// 	fmt.Println(sym, "got")
	// 	tempArr = append(tempArr, tempData)
	// 	time.Sleep(1 * time.Second)
	// }
	// SymbolStore = tempArr
	// thisPrice := calculatePrice(SymbolStore)
	// tempComplete := CompleteData{Worth: thisPrice, Errors: setErrors(), MaxWorth: setMax(thisPrice), MinWorth: setMin(thisPrice), Stocks: SymbolStore}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(TempComplete)
}

// calculate price of owned tickers
func calculatePrice(temp []SymbolData) float64 {
	var worth float64
	for _, s := range temp {
		worth += s.Price * float64(symbols[s.Symbol])
	}
	return parseFloat(worth)
}

// set errors if didnt recieve enough ticker data
func setErrors() int {
	if len(SymbolStore) != len(symbols) {
		return 1
	}
	return 0
}

// set max value if needed
func setMax(temp float64) float64 {
	if temp > Max || Max == 0.0 {
		fmt.Println("new Max", temp)
		Max = temp
	}
	return Max
}

// set min value if needed
func setMin(temp float64) float64 {
	if temp < Min || Min == 0.0 {
		fmt.Println("new Min", temp)
		Min = temp
	}
	return Min
}

// cutoff after 2nd decimal place
func parseFloat(temp float64) float64 {
	temp, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", temp), 64)
	return temp
}

// build the URL string with token
func RequestBuilder(symbol string) string {
	//fmt.Println(_remoteUrl + ":" + _remotePort + "/raw/" + symbol + "?token=" + _token)
	return _remoteUrl + ":" + _remotePort + "/raw/" + symbol + "?token=" + _token
}

// http request to data server
func RequestData(ticker string) []byte {
	resp, e := http.Get(ticker)
	if e != nil {
		fmt.Println("no data retrieved")
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return body
}
