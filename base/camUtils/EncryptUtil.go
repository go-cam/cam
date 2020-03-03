package camUtils

import (
	"crypto/md5"
	"fmt"
)

type EncryptUtil struct {
}

var Encrypt = new(EncryptUtil)

func (util *EncryptUtil) Md5(data []byte) string {
	return fmt.Sprintf("%x", md5.Sum(data))
}
