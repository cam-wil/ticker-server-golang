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

func Complete(w http.ResponseWriter, r *http.Request) {
	var tempData TickerData
	var tempArr []TickerData
	for sym, _ := range symbols {
		rd := RequestData(RequestBuilder(sym))
		err := json.Unmarshal(rd, &tempData)
		if err != nil {
			log.Fatal("failed to unmarshal", string(rd))
		}
		tempArr = append(tempArr, tempData)
		fmt.Println(tempData)
	}
	tempWorth := RoundMoney(CalculatePrice(tempArr))
	CalcMin(tempWorth)
	CalcMax(tempWorth)

	tempComplete := CompleteData{Worth: tempWorth, MaxWorth: _max, MinWorth: _min, Stocks: tempArr}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tempComplete)
}

// string builder to make URL string
func RequestBuilder(symbol string) string {
	return _remoteUrl + ":" + _remotePort + "/raw/" + symbol + "?token=" + _token
}

// get data []byte from just a stock ticker string
func RequestData(ticker string) []byte {
	resp, e := http.Get(ticker)
	if e != nil {
		fmt.Println("no data retrieved")
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return body
}
