package util

import (
	"math/rand"
	"strings"
	"time"
)

const (
	alphabet       = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	alphabetLength = len(alphabet)
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func GenerateRandomInt(min int32, max int32) int32 {
	return min + rand.Int31n(max-min+1)
}

func GenerateRandomString(length int) string {
	var stringsBuilder strings.Builder

	for i := 0; i < length; i++ {
		randomChar := generateRandomChar()
		stringsBuilder.WriteByte(randomChar)
	}

	return stringsBuilder.String()
}

func generateRandomChar() uint8 {
	return alphabet[rand.Intn(alphabetLength)]
}
