package utils

import (
	"math/rand"
	"time"

	"github.com/oklog/ulid/v2"
)

// 現在時刻をseedにしてulidを返却する
func CreateUlid() (ulid.ULID, error) {
	entropy := rand.New(rand.NewSource(time.Now().UnixNano()))
	ms := ulid.Timestamp(time.Now())
	return ulid.New(ms, entropy)
}
