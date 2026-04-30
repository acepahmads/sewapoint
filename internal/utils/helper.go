package utils

import (
	"time"
)

// format tanggal ke string
func FormatDate(t time.Time) string {
	return t.Format("2006-01-02")
}

// parse string ke date
func ParseDate(dateStr string) (time.Time, error) {
	return time.Parse("2006-01-02", dateStr)
}

// check empty string
func IsEmpty(s string) bool {
	return len(s) == 0
}
