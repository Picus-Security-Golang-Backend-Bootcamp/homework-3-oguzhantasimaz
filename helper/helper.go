package helper

import (
	"crypto/rand"
	"log"
	"math/big"
	"strconv"
	"strings"
)

//random string generator with cryptorand
func RandomString(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		value, err := rand.Int(rand.Reader, big.NewInt(int64(len(letterRunes))))
		if err != nil {
			log.Fatal(err)
		}
		b[i] = letterRunes[value.Int64()]
	}
	return string(b)
}

//random int generator with cryptorand
func RandomInt(max int64) int {
	value, err := rand.Int(rand.Reader, big.NewInt(max))
	if err != nil {
		log.Fatal(err)
	}
	return int(value.Int64())
}

//StrContains checks if a string contains a substring with case insensitive
func StrContains(s, substr string) bool {
	return strings.Contains(strings.ToLower(s), strings.ToLower(substr))
}

//IsInt checks if a string is an integer
func IsInt(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

//String to int conversion
func StringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return i
}
