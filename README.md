# Cam 
Http And Socket Framework

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/go-cam/cam?color=red)
[![GitHub tag (latest by date)](https://img.shields.io/github/v/tag/go-cam/cam)](https://github.com/go-cam/cam/tags)
![GitHub last commit](https://img.shields.io/github/last-commit/go-cam/cam)
[![GoDoc](https://godoc.org/github.com/go-cam/cam?status.svg)](https://godoc.org/github.com/go-cam/cam)
[![TODOs](https://badgen.net/https/api.tickgit.com/badgen/github.com/go-cam/cam)](https://www.tickgit.com/browse?repo=github.com/go-cam/cam)

Cam is a personal open source framework. It's goal is not to be lightweight, but to be a large framework of multi-functional integration.

Cam may not be suitable for small projects. It may be more suitable for medium and large-scale projects under multiple modules, at least this is the original intention of its development


# Contents

- [Cam](#cam)
  - [Framework structure](#framework-structure)
  - Start with template. Jump to [cam-template](https://github.com/go-cam/cam-template)
  - [Easy start](#easy-start) 
    - If you don't want cumbersome configuration. You can try this. But most documents are based on the the [cam-template](https://github.com/go-cam/cam-template)
  - [Template struct](#template-struct)
  - [Environment support](#environment-support)
    - Jump to [XORM](https://xorm.io)
  - [Components](#Components)
  - [Examples](#examples)
    - [.Env file](#env-file)
    - [Upload file](#upload-file)
    - [Validation](#validation)
    - [Middleware](#middleware)
    - [Grpc](#grpc)
- [Contact](#Contact)

## Framework structure

![Framework structure](https://raw.githubusercontent.com/go-cam/cam-doc/master/image/framework_uml.jpg)

## Easy start

### 1. Create two file:

`go.mod`
```go
module app

go 1.14

require (
	github.com/go-cam/cam v0.4.4
)
```

`main.go`
```go
package main

import (
    "github.com/go-cam/cam"
    "github.com/go-cam/cam/plugin/camRouter"
)

func main() {
	cam.RegisterController(&HelloController{})
	cam.RunDefault()
}

type HelloController struct {
	camRouter.Controller
}

func (ctrl *HelloController) Cam() {
	ctrl.SetResponse([]byte("cam is word."))
}
```

### 2. Build and run

build

    go build main.go
    
run

    ./main


Open the browser and open link: http://127.0.0.1:20200/hello/cam

## Template struct: 
    The document is explained based on this directory
 
```text
.
|-- .docker                 // docker file
|-- common                  // common module directory
  |-- config                
    |-- app.go              // common module config
  |-- templates
    |-- xorm                // xorm generate orm files's template
      |-- config
      |-- struct.go.tpl
|-- server                  // server module directory
  |-- config
    |-- app.go
    |-- bootstrap.go        // app bootstrap file
  |-- controllers           // controllers directory
    |-- HelloController.go
  |-- main.go               // entry of server module
|-- .gitignore
|-- cam.go                  // command line tools. you can build and execute it
|-- go.mod
|-- go.sum
|-- LICENSE
|-- README.md
``` 

## Environment support

- System
  - window
  - linux
- Database source (Relational database. Supports all [XORM](https://xorm.io) supported engines)
  - [Mysql5.*](https://github.com/mysql/mysql-server/tree/5.7) / [Mysql8.*](https://github.com/mysql/mysql-server) / [Mariadb](https://github.com/MariaDB/server) / [Tidb](https://github.com/pingcap/tidb)
    - [github.com/go-sql-driver/mysql](https://github.com/go-sql-driver/mysql)
    - [github.com/ziutek/mymysql/godrv](https://github.com/ziutek/mymysql/godrv)
  - [Postgres](https://github.com/postgres/postgres) / [Cockroach](https://github.com/cockroachdb/cockroach)
    - [github.com/lib/pq](https://github.com/lib/pq)
  - [SQLite](https://sqlite.org)
    - [github.com/mattn/go-sqlite3](https://github.com/mattn/go-sqlite3)
  - MsSql
    - [github.com/denisenkom/go-mssqldb](https://github.com/denisenkom/go-mssqldb)
  - Oracle
    - [github.com/mattn/go-oci8](https://github.com/mattn/go-oci8) (experiment)
- Cache engine
  - File
  - Redis

## Components

Components are the functions provided by the framework. Such as: http, websocket, cache, database, console, etc. It is a package of individual functions. 

You can also package components according to your requirements, as long as you implement camBase.ComponentInterface And corresponding configuration implementation camBase.ComponentConfigInterface that will do
  
## Examples

### .Env file
.env file must be in the same directory as the executable file.

It is recommended to create .env file in this directory: `./server/.env`

./server/.env:
```text
DB_USERNAME = root
DB_PASSWORD = 123456
```

use in code:
```go
func test() {
  username := cam.App.GetEnv("DB_USERNAME") // username = "root"
  password := cam.App.GetEnv("DB_PASSWORD") // password = "123456"
  fmt.println(username + " " + password) // output: "root 123456"
}
```

### Upload file

Example:

`FileController.go`:
```go
import (
	"github.com/go-cam/cam"
	"github.com/go-cam/cam/base/camUtils"
)

type FileController struct {
	cam.HttpController
}

func (ctrl *FileController) Upload() {
	uploadFile := ctrl.GetFile("file")
	if uploadFile == nil {
		cam.App.Fatal("FileController.Upload", "no upload file")
		return
	}

	absFilename := camUtils.File.GetRunPath() + "/runtime/upload/tmp.jpg"
	err := uploadFile.Save(absFilename)
	if err != nil {
		cam.App.Fatal("FileController.Upload", err.Error())
		return
	}

	cam.App.Trace("FileController.Upload", absFilename)
}
```

Then
post file to `http://.../file/upload`

### Validation

Example:
```go
package valid

import (
    "github.com/go-cam/cam"
    "github.com/go-cam/cam/base/camBase"
    "github.com/go-cam/cam/base/camStructs"
)

type User struct {
	Email   string
	MyEmail Email
}

type Email string

func (user *User) Rules() []camBase.RuleInterface {
	return []camBase.RuleInterface{
		camStructs.NewRule([]string{"Email", "MyEmail"}, cam.Rule.Email, cam.Rule.Length(0, 100)),
	}
}

func init() {
	user := new(User)
	user.Email = "123123"
	user.MyEmail = "123@123.com"
	firstErr, _ := cam.Valid(user)
	if firstErr != nil {
		panic(firstErr)
	}
}
```

### Middleware

Support Component: HttpComponent, WebsocketComponent, SocketComponent

add in ComponentConfig
```go
package config

import (
	"github.com/go-cam/cam"
	"github.com/go-cam/cam/base/camBase"
    "github.com/go-cam/cam/component/camHttp"
)


func httpServer() camBase.ComponentConfigInterface {
	config := camHttp.NewHttpComponentConfig(20000)
	config.Register(&controllers.TestController{}) 
	// Add middleware
	config.AddMiddleware("", &AMiddleware{}) // All route will use this Middleware
	return config
}

type AMiddleware struct {
}

func (mid *AMiddleware) Handler(ctx camBase.ContextInterface, next camBase.NextHandler) []byte {
	cam.Debug("AMiddleware", "before")
	res := next()
	cam.Debug("AMiddleware", "after")
	return res
}
```

### Grpc

#### Grpc Server
```

```

#### Grpc Client

# Contact

| Type | Content |
| ----- | ----- |
| QQ Group | 740731115 |
| Email | 1105412681@qq.com |
