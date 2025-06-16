package helpers

import (
	"crypto/md5"
	"encoding/hex"
	"time"

	"github.com/yeungon/discordbot/internal/config"
)

func GenerateHash() string {
	loc, _ := time.LoadLocation("Asia/Ho_Chi_Minh")
	time.Local = loc
	prefix := config.Get().SECRET_FIRST
	suffix := config.Get().SECRET_SECOND
	currentTime := time.Now().In(loc)
	dateString := currentTime.Format("212006") // dmyyyy format
	input := prefix + dateString + suffix
	hash := md5.Sum([]byte(input))
	return hex.EncodeToString(hash[:])
}
