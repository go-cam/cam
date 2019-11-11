package configs

import "strings"

// 数据库配置
type Database struct {
	BaseConfig

	// 驱动名字
	DriverName string
	// 地址
	Host string
	// 端口
	Port string
	// 数据库名字
	Name string
	// 用户名
	Username string
	// 密码
	Password string

	// Database file storage path. Default is: /[path to run dir]/database
	DBFileDir string
}

// 设置 migrate 路径
func (config *Database) SetDBFileDir(migrateDir string) *Database {
	config.DBFileDir = strings.Replace(migrateDir, "\\", "/", -1)
	return config
}

// get migrations's dir
func (config *Database) GetMigrateDir() string {
	return config.DBFileDir + "/migrations"
}
