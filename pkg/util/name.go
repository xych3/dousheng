package util

import (
	"strconv"
	"time"
)

func GetFileName(userID int64, title string) string {
	return strconv.FormatInt(userID, 10) + title + strconv.FormatInt(time.Now().Unix(), 10)
}
