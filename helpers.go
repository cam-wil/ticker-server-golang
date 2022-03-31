package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// environment variable setup
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
