package util

import "time"

// UnixSecondNow 返回当前的 Unix 秒时间
func UnixSecondNow() uint32 {
	return uint32(time.Now().Unix())
}
