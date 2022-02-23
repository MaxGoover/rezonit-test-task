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
func GenerateRandomInt(min int32, max int32) int32 {
	return min + rand.Int31n(max-min+1)
}

// GenerateRandomString generates a random string of length
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
