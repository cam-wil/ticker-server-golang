package main

import (
	"fmt"
	"time"
)

var _token string
var _listenPort string
var _remoteUrl string
var _remotePort string

// temporary store of what data to get, will go away after database introduced
var symbols = make(map[string]float64)

var SymbolStore []SymbolData

var Max float64
var Min float64
var TempComplete CompleteData

func main() {
	GetEnvironmentVariables()
	SetEnvironmentVariables()

	//TODO: add database return from request
	symbols["NFLX"] = 3
	symbols["MSFT"] = 2

	ticker := time.NewTicker(5 * time.Minute)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				if goodTime() {
					Timed()
				} else {
					fmt.Println("out of date / time range")
				}

			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

	HandleRoutes()

}

func goodTime() bool {
	date := time.Now().Weekday()
	hour := time.Now().Hour()
	if date == 0 || date == 6 {
		return false
	}
	if hour > 16 || hour < 8 {
		return false
	}
	return true
}
