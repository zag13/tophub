package timez

import "time"

// UnixSecondNow 返回当前的 Unix 秒时间
func UnixSecondNow() int64 {
	return time.Now().Unix()
}
