package urlshortener

import (
	"crypto/sha1"
	"encoding/hex"
)

const (
	DOMAIN string = ".link.t02smith.com"
)

func Shorten(url string) string {
	hash := Hash(url)
	return hex.EncodeToString(hash)[:6] + DOMAIN
}

func Hash(s string) []byte {
	hasher := sha1.New()
	hasher.Write([]byte(s))
	return hasher.Sum(nil)
}
