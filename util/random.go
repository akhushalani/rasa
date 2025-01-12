package util

import (
	"fmt"
	"time"

	"golang.org/x/exp/rand"
)

func init() {
	// Seed the random number generator
	rand.Seed(uint64(time.Now().UnixNano()))
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func RandomEmail() string {
	return fmt.Sprintf("%s@gmail.com", RandomString(10))
}
