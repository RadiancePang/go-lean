package tools

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	kAlphaNum = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	layout    = "20060102150405"
)

func RandString32() string {

	now := time.Now()
	nanoStr := fmt.Sprintf("%s%d", now.Format(layout), now.Nanosecond()/1000)
	if len(nanoStr) >= 32 {
		return nanoStr[0:32]
	}
	b := make([]byte, 32-len(nanoStr))
	for i := range b {
		b[i] = kAlphaNum[rand.Int63()%int64(len(kAlphaNum))]
	}
	return nanoStr + string(b)
}
