package utils

import "time"

// UnixTsFormat 格式化UNIX时间戳
func UnixTsFormat(ts int64) string {
	return time.Unix(ts, 0).Format("2006-01-02 15:04:05")
}