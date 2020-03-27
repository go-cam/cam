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
  - [Getting start](#getting-start)
  - [Template struct](#template-struct)
  - [Environment support](#environment-support)
  - [Examples](#examples)
    - [.Env file](#env-file)
    - [Upload file](#upload-file)
- [Future planning](#future-planning)

## Getting start

### 1. Clone cam-template from github

    git clone --depth=1 https://github.com/go-cam/cam-template.git -b v0.3.0

### 2. Rename folder to your project name

    mv cam-template my-project
    
### 3. Update dependency Library

    cd my-project
    go mod tidy

### 4. Build and run server module

    cd server
    go build main.go
    ./main

### 5. Check whether it runs successfully

Open the browser and visit: http://127.0.0.1:8800/hello

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
```text
username := cam.App.GetEnv("DB_USERNAME") // username = "root"
password := cam.App.GetEnv("DB_PASSWORD") // password = "123456"
fmt.println(username + " " + password) // output: "root 123456"
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

# Future planning
- v0.4.0
  - Add socket component
  - Add Struct Verification tag
