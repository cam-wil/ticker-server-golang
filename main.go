package main

var _token string
var _listenPort string
var _remoteUrl string
var _remotePort string
var symbols = make(map[string]int)
var _min float64
var _max float64

func main() {
	GetEnvironmentVariables()
	SetEnvironmentVariables()

	//TODO: add database return from request
	symbols["NFLX"] = 3
	symbols["MSFT"] = 3

	HandleRoutes()
}
