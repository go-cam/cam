[![GitHub tag (latest by date)](https://img.shields.io/github/v/tag/go-cam/cam)](https://github.com/go-cam/cam/tags)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/go-cam/cam?color=red)
![GitHub last commit](https://img.shields.io/github/last-commit/go-cam/cam)

# cin

http 和 websocket 的服务端框架。

灵感来源于 yii2 高级模板。

# 开始使用

编辑文件:  `main.go`
```go
package main

import (
    "github.com/go-cam/cam"
    "github.com/go-cam/cam/base"
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

运行命令生成依赖
> go mod tidy

编译代码
> go build main.go

运行二进制文件

- windows
> ./main.exe

- linux
> ./main


# 数据库迁移（Migrations）
编译文件后，运行一下命令生成数据库迁移文件
> ./main.exe migrate/create [filename]
