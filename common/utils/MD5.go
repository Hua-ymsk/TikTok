package utils

import (
	"crypto/md5"
	"encoding/hex"
)

// md5加密
func Md5(src string, salt string) string {
	m := md5.New()
	s := []byte(salt)
	m.Write(s)           //加盐
	m.Write([]byte(src)) //加密
	res := hex.EncodeToString(m.Sum(nil))
	return res
}
