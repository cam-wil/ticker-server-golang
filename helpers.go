package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"strings"
)

func GetEnvironmentVariables() {
	fileData, err := ioutil.ReadFile(".env")
	if err != nil {
		log.Fatal("env file not found")
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

func SetEnvironmentVariables() {
	_token = os.Getenv("TOKEN")
	_listenPort = os.Getenv("LISTENPORT")
	_remoteUrl = os.Getenv("REMOTEURL")
	_remotePort = os.Getenv("REMOTEPORT")
	fmt.Println(_token)
	fmt.Println(_listenPort)
	fmt.Println(_remoteUrl)
	fmt.Println(_remotePort)
}

func CalculatePrice(temp []TickerData) float64 {
	var worth float64
	for _, s := range temp {
		worth += s.Price * float64(symbols[s.Symbol])
	}
	return worth
}

func CalcMax(t float64) {
	if _max == 0 || t > _max {
		_max = t
	}
}

func CalcMin(t float64) {
	if _min == 0 || t < _min {
		_min = t
	}
}

func RoundMoney(t float64) float64 {
	return math.Round(t*100) / 100
}
