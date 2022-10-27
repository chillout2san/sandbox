package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

// 引数に受け取った文字列をハッシュ化して返却する。
func CreateHash(value string) string {
	cs := sha256.Sum256([]byte(value))
	h := hex.EncodeToString(cs[:])
	return h
}
