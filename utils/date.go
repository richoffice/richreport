package utils

import "time"

//2022-05-09 08:02:00 +0000 UTC
func ParseDate(dataStr string) (time.Time, error) {
	return time.Parse("2006-01-02 15:04:05 -0700 UTC", dataStr)
}
