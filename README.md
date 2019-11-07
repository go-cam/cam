[![GitHub tag (latest by date)](https://img.shields.io/github/v/tag/cinling/cin)](https://github.com/cinling/cin/tags)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/cinling/cin?color=red)
![GitHub last commit](https://img.shields.io/github/last-commit/cinling/cin)

# cin
The http and websocket server's framework.

Inspiration come form yii2.

# Getting started

main.go
```go
package main

import (
    "github.com/cinling/cin"
    "github.com/cinling/cin/base"
)

func main() {
	config := cin.NewConfig()
    config.ComponentDict = map[string]base.ConfigComponentInterface{
        "ws":      cin.NewConfigWebsocketServer(24600),
        "http":    cin.NewConfigHttpServer(24601).SetSessionName("test"),
        "db":      cin.NewConfigDatabase("mysql", "127.0.0.1", "3306", "cin", "root", "root"),
        "console": cin.NewConfigConsole(),
    }
    cin.App.AddConfig(config)
    cin.App.Run()
}
```

Then run go module to download dependency. 
```
go mod tidy
```