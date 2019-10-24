package main

import (
	"cin"
	"cin/test/backend/config"
)

func main() {
	config.LoadConfig()
	cin.App.Run()
}
