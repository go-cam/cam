package utils

import (
	"os"
	"path/filepath"
)

// 文件工具
type fileUtil struct {

}

var File = new(fileUtil)

// 获取运行程序所在的路径
func (util *fileUtil) GetRunPath() string {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	return dir
}

// 判断文件是否存在
func (util *fileUtil) Exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

// 创建目录
func (util *fileUtil) Mkdir(path string) error {
	if !util.Exists(path) {
		err := os.MkdirAll(path, os.ModePerm)
		return err
	}
	return nil
}