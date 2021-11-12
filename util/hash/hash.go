package hash

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"io/ioutil"

	"golang.org/x/crypto/bcrypt"
)

// Make a bcrypt password from a plaintext password
func Make(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), 10)
}

// Check determines if a password matches a hash
func Check(hash string, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}

// Sha1 returns a sha1 hash of a string
func Sha1(str string) string {
	hasher := sha1.New()
	hasher.Write([]byte(str))
	return hex.EncodeToString(hasher.Sum(nil))
}

// Md5 returns a md5 hash of a string
func Md5(str string) string {
	hasher := md5.New()
	hasher.Write([]byte(str))
	return hex.EncodeToString(hasher.Sum(nil))
}

// HmacSha256 returns a hmac sha256 hash of a string
func HmacSha256(str string, key string) string {
	hasher := hmac.New(sha256.New, []byte(key))
	hasher.Write([]byte(str))
	return hex.EncodeToString(hasher.Sum(nil))
}

// FileMd5 returns a md5 hash of a file
func FileMd5(path string) (string, error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	hasher := md5.New()
	hasher.Write([]byte(file))
	return hex.EncodeToString(hasher.Sum(nil)), nil
}
