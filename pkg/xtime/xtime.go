package xtime

import "time"

var TimeLocation *time.Location

func init() {
	var err error
	TimeLocation, err = time.LoadLocation("Asia/Shanghai")
	if err != nil {
		panic(err)
	}
}

func StringTime(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

func StringTimeLocateShanghai(t time.Time) string {
	return t.In(TimeLocation).Format("2006-01-02 15:04:05")
}
