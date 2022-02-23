package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const alphabetLength = len(alphabet)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// GenerateRandomInt generates a random integer between min and max
func GenerateRandomInt(min int64, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// GenerateRandomString generates a random string of length lengthString
func GenerateRandomString(stringLength int) string {
	var stringsBuilder strings.Builder

	for i := 0; i < stringLength; i++ {
		randomChar := generateRandomChar()
		stringsBuilder.WriteByte(randomChar)
	}

	return stringsBuilder.String()
}

func generateRandomChar() uint8 {
	return alphabet[rand.Intn(alphabetLength)]
}
