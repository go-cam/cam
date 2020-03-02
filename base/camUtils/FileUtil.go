package camUtils

import (
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

// file util
type FileUtil struct {
}

var File = new(FileUtil)

// Get the path where the program is running
func (util *FileUtil) GetRunPath() string {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	return dir
}

// check whether file exists
func (util *FileUtil) Exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

// make dir
// create dir
func (util *FileUtil) Mkdir(path string) error {
	if !util.Exists(path) {
		err := os.MkdirAll(path, os.ModePerm)
		return err
	}
	return nil
}

// read all content from file
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

// write all content to file
func (util *FileUtil) WriteFile(filename string, content []byte) error {
	return ioutil.WriteFile(filename, content, 0644)
}

// append content end of the file
func (util *FileUtil) AppendFile(filename string, content []byte) error {
	if !util.Exists(filename) {
		return util.WriteFile(filename, content)
	}

	file, err := os.OpenFile(filename, os.O_WRONLY, 0644)
	defer func() {
		_ = file.Close()
	}()

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

// delete file
func (util *FileUtil) DeleteFile(filename string) error {
	return os.Remove(filename)
}

// View all files in the folder
// dir:			absolute path
// withDir:		Whether the returned result contains folders
func (util *FileUtil) ScanDir(dir string, withDir bool) ([]os.FileInfo, error) {
	return ioutil.ReadDir(dir)
}

// get parent dir
func (util *FileUtil) Dir(dir string) string {
	return filepath.Dir(dir)
}

// get file size in bytes. only support file
func (util *FileUtil) Size(filename string) int64 {
	fileInfo, err := os.Stat(filename)
	if err != nil {
		return 0
	}
	if fileInfo.IsDir() {
		return 0
	}

	size := fileInfo.Size()
	return size
}

// rename file or dir
func (util *FileUtil) Rename(oldPath, newPath string) error {
	return os.Rename(oldPath, newPath)
}
