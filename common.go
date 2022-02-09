package util

import (
	"container/list"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"path"
	"reflect"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"sync/atomic"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

type TimeLayoutType string

const (
	FORMAT_LAYOUT_DATE       TimeLayoutType = "2006-01-02 15:04:05"
	FORMAT_LAYOUT_DATE_SLASH TimeLayoutType = "2006/01/02 15:04:05"
	FORMAT_LAYOUT_DAY        TimeLayoutType = "2006-01-02"
	FORMAT_LAYOUT_DAY_SLASH  TimeLayoutType = "2006/01/02"
	FORMAT_LAYOUT_TIME       TimeLayoutType = "15:04:05"
)

type OS_TYPE int32
const (
	OS_TYPE_UNKNOWN OS_TYPE = iota
	OS_TYPE_LINUX
	OS_TYPE_WINDOWS
	OS_TYPE_DARWIN
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
	rand.Seed(time.Now().UnixNano())
}

func AtomicAdd(addr *int32) int32 {
	atomic.CompareAndSwapInt32(addr, GetInt32Max(), 1)
	return atomic.AddInt32(addr, 1)
}

func StructToJson(data interface{}) string {
	if bytes, err := json.Marshal(data); err != nil {
		return ""
	} else {
		str := string(bytes)
		return str
	}
}

func StructToJsonIndent(data interface{}) string {
	if bytes, err := json.MarshalIndent(data, "", " "); err != nil {
		return ""
	} else {
		str := string(bytes)
		return str
	}
}

func JsonToStruct(jsonStr string, data interface{}) (err error) {
	err = json.Unmarshal([]byte(jsonStr), data)
	return
}

func StructToJsonWithError(data interface{}) (jsonStr string, err error) {
	var bytes []byte
	if bytes, err = json.Marshal(data); err != nil {
		return
	}
	jsonStr = string(bytes)
	return
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

func GetInt32Max() int32 {
	return 0x7fffffff
}

//ReplaceEscapeChar 去掉字符串里包含的转义字符
func ReplaceEscapeChar(src string) (des string) {
	des = src
	for k, v := range escapeCharMap {
		des = strings.Replace(des, k, v, -1)
	}
	return
}

func CopyInt32Array(src []int32) []int32 {
	ret := []int32{}
	for _, v := range src {
		ret = append(ret, v)
	}
	return ret
}

func pathExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}

	return false, err
}

func Dump(format string, args ...interface{}) {
	logPath := "../log"
	if exists, _ := pathExist(logPath); !exists {
		os.MkdirAll(logPath, os.ModePerm)
	}

	name := path.Join(logPath, "dump_"+time.Now().Format("2006-01-02_15_04")+".log")
	file, _ := os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModeAppend)
	fmt.Fprintln(file, fmt.Sprintf(format, args...))
	file.Close()
}

func FormatMoney(money int) string {
	p := message.NewPrinter(language.English)
	return p.Sprintf("%d", money)
}

//ToString 效率不高，频繁调用的话用其他实现方法
func ToString(num interface{}) (ret string) {
	return fmt.Sprintf("%v", num)
}

func MD5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func GenUUID() string {
	_uuid, err := uuid.NewRandom()
	if err != nil {
		return MD5(fmt.Sprintf("%d", time.Now().UnixNano()))
	}
	return _uuid.String()
}

func ReplaceXMLMark(str string) (ret string) {
	reg := regexp.MustCompile(`<.*?>`)
	ret = reg.ReplaceAllString(str, "")
	reg = regexp.MustCompile(`\\u003c.*?\\u003e`)
	ret = reg.ReplaceAllString(ret, "")
	return
}

func StringToInt64(str string) int64 {
	v, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0
	}
	return v
}

func Int64ToString(num int64) string {
	return strconv.FormatInt(num, 10)
}

func Float64ToString(src float64) string {
	return strconv.FormatFloat(src, 'f', -1, 64)
}

func FixIllegalNick(jsonStr string) string {
	reg := regexp.MustCompile(`"NickName":"(.*?)"[,}]`)
	result := reg.FindStringSubmatch(jsonStr)
	if len(result) > 1 {
		newStr := strings.Replace(result[1], "\"", `\"`, -1)
		jsonStr = strings.ReplaceAll(jsonStr, result[1], newStr)
	}
	return jsonStr
}

func GetOSType() OS_TYPE {
	switch runtime.GOOS {
	case "linux":
		return OS_TYPE_LINUX
	case "windows":
		return OS_TYPE_WINDOWS
	case "darwin":
		return OS_TYPE_DARWIN
	default:
		return OS_TYPE_UNKNOWN
	}
}

//只能确定1开头,第二位3-9,没有其他确定依据
func CheckChinaPhoneNumValid(phoneNum string) bool {
	result, _ := regexp.MatchString(`^(1[3-9][0-9]{9})$`, phoneNum)
	if result {
		return true
	} else {
		return false
	}
}