package pkg

import (
	"crypto/rand"
	"encoding/hex"
)

// 生成字符串(32位）
func GenerateRandomString() string {
	bytes := make([]byte, 32/2)
	_, _ = rand.Read(bytes)
	return hex.EncodeToString(bytes)
}
