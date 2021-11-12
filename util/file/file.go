package file

import "os"

// Exist checks if a file or directory exists.
func Exist(name string) bool {
	if _, err := os.Stat(name); os.IsNotExist(err) {
		return false
	}
	return true
}

// MkdirAll creates a directory path recursively.
func MkdirAll(path string) error {
	if err := os.MkdirAll(path, 0755); err != nil {
		return err
	}
	return nil
}
