package gofunc

import "time"

// timestamp.
// unit is seconds.
func Ts() int64 {
	return time.Now().Unix()
}

// timestamp.
// unit is millisecond.
func TsMs() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

