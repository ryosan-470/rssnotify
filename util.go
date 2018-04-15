package main

import (
	"time"
)

// ToTime converts argument to time.Time
func ToTime(t string) time.Time {
	// "Fri, 13 Apr 2018 09:06:00 GMT"
	time, _ := time.Parse(time.RFC1123, t)
	return time
}

// IsUpdated は updated が 現在時刻から interval (min) 以内に更新されたかどうかを判定する
func IsUpdated(interval int, updated, now time.Time) bool {
	t := time.Duration(interval) * time.Minute
	newUpdated := updated.Add(t)
	return newUpdated.After(now)
}
