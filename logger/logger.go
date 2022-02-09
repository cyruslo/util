package logger

import (
	"fmt"
	"log"
	"os"

	//	"testing"
	"runtime"
	"time"

	ct "github.com/daviddengcn/go-colortext"
)

var (
	logger           *log.Logger
	start            = false
	lastCreateLogSec int64
	strlogPathName   string
	fileSize         = int64(100 * 1024 * 1024)
	fileName         string
)

func getLogName(name string) string {
	t := time.Now()
	logname := fmt.Sprintf("%s_%04d_%02d_%02d_%02d_%02d_%02d_%d.log",
		name,
		t.Year(),
		t.Month(),
		t.Day(),
		t.Hour(),
		t.Minute(),
		t.Second(),
		os.Getpid())

	return logname
}

//获取文本的大小
func getFileSize(fileName string) int64 {
	fInfo, err := os.Stat(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			OnStartup(strlogPathName) //文件不见了? 重建
		}
		log.Printf("getfilesize err %v", err)

		return 0
	}
	
	return fInfo.Size()
}

//log相关
func Info(format string, args ...interface{}) {
	str := fmt.Sprintf(format, args...)

	ct.Foreground(ct.Green, false)
	log.Print(str)
	ct.ResetColor()

	//文本记录
	pc, _, line, _ := runtime.Caller(1)
	f := runtime.FuncForPC(pc)
	str = fmt.Sprintf("%s %d Info: %s", f.Name(), line, str)
	logger.Print(str)

	cursec := getCurSec()
	if diffTimeDay(lastCreateLogSec, cursec) == true || getFileSize(fileName) >= fileSize {
		OnStartup(strlogPathName)
	}
}

func Warning(format string, args ...interface{}) {
	str := fmt.Sprintf(format, args...)

	ct.Foreground(ct.Yellow, false)
	log.Print(str)
	ct.ResetColor()

	//文本记录
	pc, _, line, _ := runtime.Caller(1)
	f := runtime.FuncForPC(pc)
	str = fmt.Sprintf("%s %d Warning: %s", f.Name(), line, str)
	logger.Print(str)

	cursec := getCurSec()
	if diffTimeDay(lastCreateLogSec, cursec) == true || getFileSize(fileName) >= fileSize {
		OnStartup(strlogPathName)
	}
}

func Error(format string, args ...interface{}) {
	str := fmt.Sprintf(format, args...)

	ct.Foreground(ct.Red, false)
	log.Print(str)
	ct.ResetColor()

	//文本记录
	pc, _, line, _ := runtime.Caller(1)
	f := runtime.FuncForPC(pc)
	str = fmt.Sprintf("%s %d Error: %s", f.Name(), line, str)
	logger.Print(str)

	cursec := getCurSec()
	if diffTimeDay(lastCreateLogSec, cursec) == true || getFileSize(fileName) >= fileSize {
		OnStartup(strlogPathName)
	}
}

//获取当前秒
func getCurSec() int64 {
	return time.Now().Unix()
}

func diffTimeDay(early int64, later int64) bool {
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

func OnStartup(strLogName string) {
	//创建目录
	strFileName := "./log_game_trace/"
	res := os.MkdirAll(strFileName, 0666)
	if res != nil {
		log.Fatalln("fail to mkdir path error(%v) !", res)
		return
	}

	nomarlLogName := fmt.Sprintf("%s%s", strFileName, getLogName(strLogName))
	file, err := os.Create(nomarlLogName)
	if err != nil {
		log.Fatalln("fail to create %s.log file!", nomarlLogName)
		return
	}

	logger = log.New(file, "", log.Ldate|log.Lmicroseconds)
	log.SetFlags(log.Lmicroseconds)

	lastCreateLogSec = getCurSec()
	strlogPathName = strLogName
	fileName = nomarlLogName
}
