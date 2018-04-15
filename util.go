package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/mmcdole/gofeed"
)

// ToTime converts argument to time.Time
func ToTime(t string) time.Time {
	// "Fri, 13 Apr 2018 09:06:00 GMT"
	time, _ := time.Parse(time.RFC1123, t)
	return time
}

// IsUpdated は updated が 現在時刻から interval (min) 以内に更新されたかどうかを判定する
func IsUpdated(interval int, updated, now time.Time) bool {
	// now >= updated
	if updated.After(now) {
		return false
	}
	t := time.Duration(interval) * time.Minute
	past := now.Add(-t)
	// past < updated < now
	fmt.Println(past)
	return updated.After(past) && updated.Before(now)
}

// FilterWithDublinCore は Itemのうち、インターバル時間以内のものだけを取り出す
func FilterWithDublinCore(items []gofeed.Item, now time.Time) []gofeed.Item {
	ret := []gofeed.Item{}
	if len(items) == 0 {
		return ret
	}

	for _, item := range items {
		date := item.DublinCoreExt.Date[0]
		t, _ := time.Parse("2006-01-02T15:04:05Z", date)
		interval, _ := strconv.ParseInt(cfg.Interval, 10, 0)
		fmt.Printf("IsUpdated(%v, %v, %v)\n", interval, t, now)
		if IsUpdated(int(interval), t, now) {
			ret = append(ret, item)
		}
	}
	return ret
}
