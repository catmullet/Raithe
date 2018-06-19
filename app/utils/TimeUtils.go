package utils

import "time"

func MakeTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func MakeTimestampTimeout(minutes int) int64 {
	return time.Now().Add(time.Duration(minutes) * time.Minute).UnixNano() / int64(time.Millisecond)
}
