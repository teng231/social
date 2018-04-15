package core

import (
	"crypto/sha1"
	"fmt"
	"time"

	"github.com/my0sot1s/social/utils"
)

const (
	secretKey = "97CGQGl4iIQ35AOWYpu_3u2S564"
)

func (c *Social) SignFileToUpload() (string, int64) {

	loc, err := time.LoadLocation("America/New_York")
	utils.ErrLog(err)
	t := time.Now().In(loc)
	timeStamp := t.Unix() * 1000
	hashString := fmt.Sprintf("timestamp=%d%s", t.Unix()*1000, secretKey)
	hash := sha1.New()
	hash.Write([]byte(hashString))
	hashCode := hash.Sum(nil)
	signature := fmt.Sprintf("%x", hashCode)
	return signature, timeStamp
}
