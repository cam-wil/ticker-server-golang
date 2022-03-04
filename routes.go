package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleRoutes() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/complete", Complete)
	log.Fatal(http.ListenAndServe(":"+_listenPort, router))
}

// routes

func Complete(w http.ResponseWriter, r *http.Request) {
	var tempData SymbolData
	var tempArr []SymbolData
	for sym, _ := range symbols {
		rd := RequestData(RequestBuilder(sym))
		err := json.Unmarshal(rd, &tempData)
		if err != nil {
			log.Fatal("failed to unmarshal", string(rd))
		}
		tempArr = append(tempArr, tempData)
		fmt.Println(tempData)
	}
	tempComplete := CompleteData{Worth: calculatePrice(tempArr), MaxWorth: 2222.22, MinWorth: 1111.11, Stocks: tempArr}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tempComplete)
}

func calculatePrice(temp []SymbolData) float64 {
	var worth float64
	for _, s := range temp {
		worth += s.Price * float64(symbols[s.Symbol])
	}
	return worth
}

func RequestBuilder(symbol string) string {
	fmt.Println(_remoteUrl + ":" + _remotePort + "/raw/" + symbol + "?token=" + _token)
	return _remoteUrl + ":" + _remotePort + "/raw/" + symbol + "?token=" + _token
}

func RequestData(ticker string) []byte {
	resp, e := http.Get(ticker)
	if e != nil {
		fmt.Println("no data retrieved")
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return body
}
