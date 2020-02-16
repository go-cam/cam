package camUtils

import (
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

// 文件工具
type FileUtil struct {
}

var File = new(FileUtil)

// 获取运行程序所在的路径
func (util *FileUtil) GetRunPath() string {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	return dir
}

// 判断文件是否存在
func (util *FileUtil) Exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

// 创建目录
func (util *FileUtil) Mkdir(path string) error {
	if !util.Exists(path) {
		err := os.MkdirAll(path, os.ModePerm)
		return err
	}
	return nil
}

// 读取文件内所有数据
func (util *FileUtil) ReadFile(filePath string) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

// 写文件
func (util *FileUtil) WriteFile(filename string, content []byte) error {
	return ioutil.WriteFile(filename, content, 0644)
}

// append content end of the file
func (util *FileUtil) AppendFile(filename string, content []byte) error {
	if !util.Exists(filename) {
		return util.WriteFile(filename, content)
	}

	file, err := os.OpenFile(filename, os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	index, err := file.Seek(0, io.SeekEnd)
	if err != nil {
		return err
	}
	_, err = file.WriteAt(content, index)
	return err
}

// 删除文件
func (util *FileUtil) DeleteFile(filename string) error {
	return os.Remove(filename)
}

// 获取目录下所有的文件列表（仅遍历，非递归）
// dir 绝对路径
// withDir 返回结果是否包含文件夹
func (util *FileUtil) ScanDir(dir string, withDir bool) ([]os.FileInfo, error) {
	return ioutil.ReadDir(dir)
}
