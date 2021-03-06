package util

import (
	"fmt"
	"time"
)

func ParseTimeStamp(date string, typo TimeLayoutType) (t int64, err error) {
	switch typo {
	case FORMAT_LAYOUT_TIME:
		date = fmt.Sprintf("%s %s", GetDay(), date)
		typo = FORMAT_LAYOUT_DATE
	case FORMAT_LAYOUT_DAY:
		date = fmt.Sprintf("%s %s", date, "00:00:00")
		typo = FORMAT_LAYOUT_DATE
	case FORMAT_LAYOUT_DAY_SLASH:
		date = fmt.Sprintf("%s %s", date, "00:00:00")
		typo = FORMAT_LAYOUT_DATE_SLASH
	}
	var stamp time.Time
	stamp, err = time.ParseInLocation(string(typo), date, time.Local)
	if err != nil {
		return
	}
	t = stamp.Unix()
	return
}

func FormatPassTime(passTime int64) string {
	if passTime < 60 {
		return fmt.Sprintf("%d秒", passTime)
	} else if passTime < 60*60 {
		return fmt.Sprintf("%.3f分钟", float64(passTime)/60)
	} else if passTime < 60*60*48 {
		//2天以内还是按小时算
		return fmt.Sprintf("%.3f小时", float64(passTime)/60/60)
	} else {
		return fmt.Sprintf("%.3f天", float64(passTime)/60/60/24)
	}
}

func GetCurHour() int32 {
	hour, _, _ := time.Now().Clock()
	return int32(hour)
}
