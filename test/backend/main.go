package main

import (
	"github.com/cinling/cin"
	"github.com/cinling/cin/test/backend/config"
)

func main() {
	config.LoadConfig()
	cin.App.Run()
}
