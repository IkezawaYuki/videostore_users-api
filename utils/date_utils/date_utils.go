package date_utils

import "time"

const (
	dateFormat = "2006-01-02 15:04:05"
)

func GetNow() time.Time{
	return time.Now()
}

func GetNowString() string{
	return GetNow().Format(dateFormat)
}