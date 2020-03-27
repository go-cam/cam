package camHttp

import (
	"github.com/go-cam/cam/base/camUtils"
	"io"
	"mime/multipart"
	"os"
	"strings"
)

// UploadFile helper
type UploadFile struct {
	File   multipart.File
	Header *multipart.FileHeader
}

// New UploadFile instance
func NewUploadFile(file multipart.File, Header *multipart.FileHeader) *UploadFile {
	uf := new(UploadFile)
	uf.File = file
	uf.Header = Header
	return uf
}

// Save file
// absFilename
func (uf *UploadFile) Save(absFilename string) error {
	dir := camUtils.File.Dir(absFilename)
	if !camUtils.File.Exists(dir) {
		err := camUtils.File.Mkdir(dir)
		if err != nil {
			return err
		}
	}

	file, err := os.Create(absFilename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, uf.File)
	return err
}

// file size
func (uf *UploadFile) Size() int64 {
	return uf.Header.Size
}

// file extension
func (uf *UploadFile) Extension() string {
	tmpArr := strings.Split(uf.Header.Filename, ".")
	length := len(tmpArr)
	if length <= 1 {
		// no extension. Example: uf.Header.Filename = "output"
		return ""
	}
	return tmpArr[length-1:][0]
}

// file suffix
func (uf *UploadFile) Suffix() string {
	return uf.Extension()
}

// filename
func (uf *UploadFile) Filename() string {
	return uf.Header.Filename
}

// Close io
func (uf *UploadFile) Close() error {
	return uf.File.Close()
}
