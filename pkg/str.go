package pkg

import (
	"crypto/rand"
	"encoding/hex"
)

// 生成字符串(32位）
func generateRandomString() (string, error) {
	bytes := make([]byte, 32/2)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
