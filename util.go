package main

import "time"

var now = time.Now()

// ToTime converts argument to time.Time
func ToTime(t string) time.Time {
	// "Fri, 13 Apr 2018 09:06:00 GMT"
	time, _ := time.Parse(time.RFC1123, t)
	return time
}
