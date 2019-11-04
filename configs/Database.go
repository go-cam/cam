package configs

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

	// 数据库迁移文件路径
	MigrateDir string
}

// 设置 migrate 路径
func (config *Database) SetMigrateDir(migrateDir string) *Database {
	config.MigrateDir = migrateDir
	return config
}