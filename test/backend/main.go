package main

import (
	"github.com/go-cam/cam"
	"test/backend/config"
)

func main() {
	config.LoadConfig()
	cam.App.Run()
}
