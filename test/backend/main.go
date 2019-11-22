package main

import (
	"github.com/cinling/cam"
	"github.com/cinling/cam/test/backend/config"
)

func main() {
	config.LoadConfig()
	cin.App.Run()
}
