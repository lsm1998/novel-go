package utils

import (
	"crypto/md5"
	"fmt"
)

func Md5(val string) string {
	data := []byte(val)
	has := md5.Sum(data)
	return fmt.Sprintf("%x", has)
}
