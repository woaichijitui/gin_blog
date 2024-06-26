package utils

import (
	"crypto/md5"
	"encoding/hex"
)

// Md5 计算文件hash值
func Md5(data []byte) string {
	// 创建一个新的MD5哈希对象
	hash := md5.New()

	// 写入数据到哈希对象
	hash.Write(data)

	// 计算哈希值
	hashInBytes := hash.Sum(nil)

	// 将字节数组转换为十六进制字符串
	hashInHex := hex.EncodeToString(hashInBytes)
	return hashInHex
}
