package util

import (
	"container/list"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"math/rand"
	"reflect"
	"strconv"
	"strings"
	"sync/atomic"
	"time"

	"golang.org/x/text/language"

	"golang.org/x/text/message"

	"github.com/bilibili/kratos/pkg/log"
)

var (
	escapeCharMap map[string]string
)

func init() {
	escapeCharMap = map[string]string{
		"\\":  "",
		"\a":  " ",
		"\b":  " ",
		"\f":  " ",
		"\n":  " ",
		"\r":  " ",
		"\t":  " ",
		"\v":  " ",
		"\\?": "?",
		"\\0": "0",
	}
}

type IWheelItem interface {
	GetWeight() int
}

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

//RandomClosed get random num [min,max]
func RandomClosed(min int, max int) int {
	if min > max {
		min, max = max, min
	}
	return rand.Intn(max-min+1) + min
}

func RandClosed(min int32, max int32) int32 {
	if min > max {
		min, max = max, min
	}
	return rand.Int31n(max-min+1) + min
}

//RandomSemiClosed get random num [min,max)
func RandomSemiClosed(min int, max int) int {
	if min == max {
		return min
	}

	if min > max {
		min, max = max, min
	}
	return rand.Intn(max-min) + min
}

func RandSemiClosed(min int32, max int32) int32 {
	if min == max {
		return min
	}

	if min > max {
		min, max = max, min
	}
	return rand.Int31n(max-min) + min
}

func RandCheck(ratio int, base int) bool {
	if base <= 0 {
		return false
	}
	//[0,base)
	randNum := rand.Intn(base)
	return ratio > randNum
}

func RandOk(ratio int32, base int32) bool {
	if base <= 0 {
		return false
	}
	//[0,base)
	return ratio > rand.Int31n(base)
}

func AtomicAdd(addr *int32) int32 {
	atomic.CompareAndSwapInt32(addr, 0x7fffffff, 1)
	return atomic.AddInt32(addr, 1)
}

func StructToJson(data interface{}) string {
	if reflect.TypeOf(data).Kind() == reflect.Ptr {
		data = reflect.ValueOf(data).Elem().Interface()
	}
	if bytes, err := json.Marshal(data); err != nil {
		return ""
	} else {
		str := string(bytes)
		return str
	}
}

func StructToJsonIndent(data interface{}) string {
	if reflect.TypeOf(data).Kind() == reflect.Ptr {
		data = reflect.ValueOf(data).Elem().Interface()
	}
	if bytes, err := json.MarshalIndent(data, "", " "); err != nil {
		return ""
	} else {
		str := string(bytes)
		return str
	}
}

func MakeError(format string, args ...interface{}) error {
	return errors.New(fmt.Sprintf(format, args...))
}

func ListForEach(list list.List, visitor func(data interface{})) {
	for iter := list.Front(); iter != nil; iter = iter.Next() {
		visitor(iter.Value)
	}
}

//Foreach 遍历容器元素执行visitor container必须是Map/Slice/Array
func ForEach(container interface{}, visitor func(k interface{}, v interface{})) {
	if !AssertContainer(container) {
		return
	}
	cValue := reflect.ValueOf(container)
	switch cValue.Type().Kind() {
	case reflect.Map:
		iter := cValue.MapRange()
		for iter.Next() {
			visitor(iter.Key().Interface(), iter.Value().Interface())
		}
	case reflect.Array, reflect.Slice:
		for i := 0; i < cValue.Len(); i++ {
			visitor(i, cValue.Index(i).Interface())
		}
	}
}

//AssertContainer 判断传入参数是否Map/Slice/Array
func AssertContainer(container interface{}) bool {
	if container == nil {
		return false
	}
	kind := reflect.ValueOf(container).Type().Kind()
	if kind != reflect.Map && kind != reflect.Slice && kind != reflect.Array {
		return false
	}
	return true
}

//AssertArray 判断传入参数是否Slice/Array
func AssertArray(container interface{}) bool {
	kind := reflect.ValueOf(container).Type().Kind()
	if kind != reflect.Slice && kind != reflect.Array {
		return false
	}
	return true
}

//GetRandomElement 使用轮盘算法随机从容器中获取一个元素 元素必须实现IWheelItem接口
func GetRandomElement(container interface{}) interface{} {
	if !AssertContainer(container) {
		return nil
	}
	cValue := reflect.ValueOf(container)
	cType := cValue.Type()
	if cValue.Len() == 0 {
		return nil
	}
	eleType := cType.Elem()
	_, succ := reflect.New(eleType).Interface().(IWheelItem)
	if !succ {
		return nil
	}
	total := 0
	ForEach(container, func(k interface{}, v interface{}) {
		total += v.(IWheelItem).GetWeight()
	})
	randNum := RandomClosed(0, total-1)
	tmpNum := 0
	var ret interface{} = nil
	ForEach(container, func(k interface{}, v interface{}) {
		tmpNum += v.(IWheelItem).GetWeight()
		if tmpNum > randNum && ret == nil {
			ret = v
		}
	})
	return ret
}

//M个里取N个
func GetNFromM(container interface{}, num int) []interface{} {
	ret := []interface{}{}
	if !AssertContainer(container) {
		return ret
	}
	cValue := reflect.ValueOf(container)
	if cValue.Len() == 0 {
		return ret
	}
	index := 0
	//蓄水池
	ForEach(container, func(k interface{}, v interface{}) {
		if index < num {
			ret = append(ret, v)
		} else {
			if RandCheck(num, index+1) {
				swapIndex := RandomClosed(0, num-1)
				ret[swapIndex] = v
			}
		}
		index++
	})
	return ret
}

//Filter 筛选Map/Slice/Array
func Filter(container interface{}, filter func(k interface{}, v interface{}) bool) []interface{} {
	ret := []interface{}{}
	if !AssertContainer(container) {
		return ret
	}
	cValue := reflect.ValueOf(container)
	if cValue.Len() == 0 {
		return ret
	}
	ForEach(container, func(k interface{}, v interface{}) {
		if filter(k, v) {
			ret = append(ret, v)
		}
	})
	return ret
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

func Shuffle(array interface{}) {
	if !AssertArray(array) {
		return
	}
	cValue := reflect.ValueOf(array)
	len := cValue.Len()
	for i := 0; i < len-1; i++ {
		randNum := RandomClosed(i, len-1)
		if randNum != i {
			tmp := cValue.Index(i).Interface()
			cValue.Index(i).Set(reflect.ValueOf(cValue.Index(randNum).Interface()))
			cValue.Index(randNum).Set(reflect.ValueOf(tmp))
		}
	}
}

func GetInt32Max() int32 {
	return 0x7fffffff
}

func ReplaceEscapeChar(src string) (des string) {
	des = src
	for k, v := range escapeCharMap {
		des = strings.Replace(des, k, v, -1)
	}
	return
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

func CopyInt32Array(src []int32) []int32 {
	ret := []int32{}
	for _, v := range src {
		ret = append(ret, v)
	}
	return ret
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

func FormatMoney(money int32) string {
	p := message.NewPrinter(language.English)
	return p.Sprintf("%d", money)
}
