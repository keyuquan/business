package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

// Md5Encrypt MD5加密
func Md5Encrypt(str string) string {
	md := md5.New()
	md.Write([]byte(str))                                   // 需要加密的字符串为 str
	cipherStr := md.Sum(nil)                                //不需要拼接额外的数据，如果约定了额外加密数据，可以在这里传递
	return fmt.Sprintf("%s", hex.EncodeToString(cipherStr)) // 输出加密结果
}

// Md5EncryptBytes MD5加密
func Md5EncryptBytes(str []byte) string {
	md := md5.New()
	md.Write([]byte(str))                                   // 需要加密的字符串为 str
	cipherStr := md.Sum(nil)                                //不需要拼接额外的数据，如果约定了额外加密数据，可以在这里传递
	return fmt.Sprintf("%s", hex.EncodeToString(cipherStr)) // 输出加密结果
}
