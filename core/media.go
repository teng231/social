package core

import (
	"crypto/sha1"
	"fmt"
	"time"
)

const (
	secretKey = "97CGQGl4iIQ35AOWYpu_3u2S564"
)

func (c *Core) SignFileToUpload() string {
	t := time.Now()
	hashString := fmt.Sprintf("timestamp=%d%s", t.Unix(), secretKey)
	hash := sha1.New()
	hash.Write([]byte(hashString))
	hashCode := hash.Sum(nil)
	signature := fmt.Sprintf("%x", hashCode)
	return signature
}
