![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/go-cam/cam?color=red)
[![GitHub tag (latest by date)](https://img.shields.io/github/v/tag/go-cam/cam)](https://github.com/go-cam/cam/tags)
![GitHub last commit](https://img.shields.io/github/last-commit/go-cam/cam)

# cin
[简体中文](https://github.com/go-cam/cam/blob/master/doc/README.zh-cn.md)

The http and websocket server's framework.

Inspiration come form yii2.

# Dependencies
	github.com/go-sql-driver/mysql v1.4.1
	github.com/go-xorm/xorm v0.7.9
	github.com/gorilla/sessions v1.2.0
	github.com/gorilla/websocket v1.4.1
	github.com/kr/pretty v0.1.0 // indirect
	github.com/tidwall/gjson v1.3.4
	golang.org/x/crypto v0.0.0-20190426145343-a29dc8fdc734 // indirect

# Getting started

Edit file:  `main.go`
```go
package main

import (
    "github.com/go-cam/cam"
    "github.com/go-cam/cam/core/base"
)

func main() {
	config := cam.NewConfig()
	config.ComponentDict = map[string]base.ConfigComponentInterface{
		"ws":      cam.NewConfigWebsocketServer(24600),
		"http":    cam.NewConfigHttpServer(24601).SetSessionName("test"),
		"db":      cam.NewConfigDatabase("mysql", "127.0.0.1", "3306", "cin", "root", "root"),
		"console": cam.NewConfigConsole(),
	 }
	 cam.App.AddConfig(config)
	 cam.App.Run()
}
```

Run go module to download dependency. 
> go mod tidy

Compiled code.
> go build main.go

Run bin file
> ./main.exe  `or` ./main


# Migrations
After Compiled code, run the following command to create migration's file.
> ./main.exe migrate/create [filename]