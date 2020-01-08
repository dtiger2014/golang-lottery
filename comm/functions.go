package comm

import (
	"golang-lottery/conf"
	"time"
)

// NowUnix : Get current timestamp
// return : int
func NowUnix() int {
	return int(time.Now().In(conf.SysTimeLocation).Unix())
}

// FormatFromUnixTime : Format unixtimestamp (int64) to date string
// if t less then 0, then return current date.
// Exp : "2020-01-02 03:04:05"
func FormatFromUnixTime(t int64) string {
	if t > 0 {
		return time.Unix(t, 0).Format(conf.SysTimeform)
	} else {
		return time.Now().Format(conf.SysTimeform)
	}
}

// FormatFromUnixTimeShort : Format unix timestamp (int64) to short date string
// if t < 0, then return current date.
// Exp: "2020-01-02"
func FormatFromUnixTimeShort(t int64) string {
	if t > 0 {
		return time.Unix(t, 0).Format(conf.SysTimeformShort)
	} else {
		return time.Now().Format(conf.SysTimeformShort)
	}
}


