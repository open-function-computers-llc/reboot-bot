package main

import (
	"errors"
	"os"
	"strconv"

	_ "github.com/joho/godotenv/autoload"
	"github.com/open-function-computers-llc/reboot-bot/server"
)

func main() {
	port, key, err := processENV()
	if err != nil {
		panic(err)
	}

	s, err := server.New(port, key)
	if err != nil {
		panic(err)
	}

	err = s.Serve()
	if err != nil {
		panic(err)
	}
}

func processENV() (int, string, error) {
	requiredENV := []string{
		"APP_PORT",
		"API_KEY",
	}

	for _, env := range requiredENV {
		check := os.Getenv(env)
		if check == "" {
			return 0, "", errors.New("missing env: " + env)
		}
	}

	webPort := os.Getenv("APP_PORT")
	portInt, err := strconv.Atoi(webPort)
	if err != nil {
		return 0, "", errors.New("Invalid APP_PORT: " + err.Error())
	}

	return portInt, os.Getenv("API_KEY"), nil
}
