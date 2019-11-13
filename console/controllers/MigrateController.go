package controllers

import (
	"bufio"
	"fmt"
	"github.com/cinling/cin/core/utils"
	"os"
	"strconv"
	"strings"
	"time"
)

// 数据库版本管理控制器
type MigrateController struct {
	baseConsoleController
}

// 创建一个数据库升级文件
func (controller *MigrateController) Create() {
	var err error

	// 生成路径
	migrateDir := controller.GetValue("migrateDir").(string)
	if utils.File.Exists(migrateDir) {
		err = utils.File.Mkdir(migrateDir)
		utils.Error.Panic(err)
	}

	// 生成id
	timestamp := time.Now().Unix()
	timestampStr := strconv.FormatInt(timestamp, 10)

	// 名字
	name := "new_migrate"
	if len(os.Args) >= 3 {
		name = os.Args[2]
	}
	upFilename := migrateDir + "/" + timestampStr + "_" + name + ".up.sql"
	downFilename := migrateDir + "/" + timestampStr + "_" + name + ".down.sql"

	fmt.Println("General filename...")
	fmt.Println("\t" + downFilename)
	fmt.Println("\t" + upFilename)
	fmt.Print("Do you want to create the following two file?[Y/N]:")
	input := bufio.NewScanner(os.Stdin)
	if !input.Scan() {
		return
	}
	str := strings.ToLower(input.Text())
	if str != "y" {
		return
	}

	if !utils.File.Exists(migrateDir) {
		err = utils.File.Mkdir(migrateDir)
		utils.Error.Panic(err)
	}

	err = utils.File.WriteFile(downFilename, []byte{})
	utils.Error.Panic(err)
	err = utils.File.WriteFile(upFilename, []byte{})
	utils.Error.Panic(err)

	fmt.Println("")
	fmt.Println("Done: migrations's files created.")
}
