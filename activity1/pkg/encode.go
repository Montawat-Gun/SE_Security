package activity1

import (
	"crypto/sha1"
	"encoding/hex"
)

var shaOne = sha1.New()

func HashSha1(word string) string {
	shaOne.Reset()
	shaOne.Write([]byte(word))
	return hex.EncodeToString(shaOne.Sum(nil))
}
