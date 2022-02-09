package util

import (
	"github.com/bilibili/kratos/pkg/log"
	"math"
	"reflect"
	"strconv"
	"time"
)

func GetTimeStamp() int64 {
	return time.Now().Unix()
}

func GetTimeStampMs() int64 {
	return time.Now().UnixNano() / (1000 * 1000)
}

func DiffTimeDay(early int64, later int64) bool {
	if early > later {
		early, later = later, early
	}

	t1 := time.Unix(early, 0)
	t2 := time.Unix(later, 0)

	if t1.Year() != t2.Year() || t1.Month() != t2.Month() || t1.Day() != t2.Day() {
		return true
	}

	return false
}

//find 筛选Map/Slice/Array
func Finder(container interface{}, filter func(v interface{}) bool) interface{} {

	if !AssertContainer(container) {
		return nil
	}

	cValue := reflect.ValueOf(container)
	if cValue.Len() == 0 {
		return nil
	}

	switch cValue.Type().Kind() {
	case reflect.Map:
		iter := cValue.MapRange()
		for iter.Next() {
			if filter(iter.Value()) {
				return iter.Value().Interface()
			}
		}
	case reflect.Array, reflect.Slice:
		for i := 0; i < cValue.Len(); i++ {
			if filter(cValue.Index(i)) {
				return cValue.Index(i).Interface()
			}
		}
	}

	return nil
}

func GetDate() string {
	return FormatDate(GetTimeStamp())
}

func GetDay() string {
	return FormatDay(GetTimeStamp())
}

func FormatDay(timestamp int64) string {
	timeTemplate := "2006-01-02"
	return time.Unix(timestamp, 0).Format(timeTemplate)
}

func FormatDate(timestamp int64) string {
	timeTemplate := "2006-01-02 15:04:05"
	return time.Unix(timestamp, 0).Format(timeTemplate)
}

//string、int、int64相互转换
//string转到int
func StrToInt(src string) int {
	//log.Printf("StrToInt:%s.", src)
	dst, err := strconv.Atoi(src)
	if err != nil {
		log.Error("str to int error(%v).", err)
		return 0
	}
	return dst
}

//string转到int64
func StrToInt64(src string) int64 {
	//	log.Print("src:%s.", src)
	dst, err := strconv.ParseInt(src, 10, 64)
	if err != nil {
		log.Error("str to int64 error(%v).", err)
		return 0
	}
	return dst
}

//int转到string
func IntToStr(src int) string {
	dst := strconv.Itoa(src)
	return dst
}

//int64转到string
func Int64ToStr(src int64) string {
	dst := strconv.FormatInt(src, 10)
	return dst
}

//string转到int32
func StrToInt32(src string) int32 {
	dst, err := strconv.Atoi(src)
	if err != nil {
		log.Error("str to int32 error(%v).", err)
		return 0
	}
	return int32(dst)
}

//int转到string
func Int32ToStr(src int32) string {
	dst := strconv.Itoa(int(src))
	return dst
}

//整数保留最高几位非零，其他位数补零
func RoundInt(num int64, keepBits int32) int64 {
	temp := num
	count := int32(0)

	for {
		count++
		temp /= 10
		if temp <= 0 {
			break
		}
	}

	if keepBits > 0 && keepBits < count {
		multi := int64(math.Pow(10.0, float64(count-keepBits)))
		if multi > 0 {
			return (num / multi) * multi
		}
	}

	return num
}