package str

import (
	"math/rand"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"
)

const (
	pool = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

var regexCamel = regexp.MustCompile("([a-z])([A-Z])")

// Len returns the length of the utf8 string.
func Len(str string) int {
	return utf8.RuneCountInString(str)
}

// Random returns a random string of the specified length.
func Random(length int) string {
	letters := make([]byte, length)

	rand.Seed(time.Now().UnixNano())
	for i := range letters {
		letters[i] = pool[rand.Intn(len(pool))]
	}
	return string(letters)
}

// ToSnakeCase converts a string to snake case.
func ToSnakeCase(str string) string {
	return strings.ToLower(regexCamel.ReplaceAllString(str, "${1}_${2}"))
}
