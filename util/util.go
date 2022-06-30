package util

import "crypto/sha1"

// Return the sha1 hash of a string
func Hash(s string) []byte {
	hasher := sha1.New()
	hasher.Write([]byte(s))
	return hasher.Sum(nil)
}
