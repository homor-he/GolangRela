package basefunc

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
	"time"
)

type MyLog struct {
	fileName string
	fileObj  *os.File
}

func (myLog *MyLog) Init() {
	var err error
	currentTime := time.Now()
	myLog.fileName = fmt.Sprintf("%d-%d-%d autotest.log", currentTime.Year(), currentTime.Month(), currentTime.Day())
	myLog.fileObj, err = os.OpenFile(myLog.fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Printf("open log file:%s fail", myLog.fileName)
		return
	}
	//log.SetOutput(myLog.fileObj)
	//log.SetFlags(log.Ltime | log.Lshortfile)
}

func (myLog *MyLog) Print(msg string) {
	var err error
	currentTime := time.Now()
	sFileName := fmt.Sprintf("%d-%d-%d autotest.log", currentTime.Year(), currentTime.Month(), currentTime.Day())
	if sFileName != myLog.fileName {
		myLog.fileName = sFileName
		myLog.fileObj, err = os.OpenFile(myLog.fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			log.Printf("open log file:%s fail", myLog.fileName)
			return
		}
		//log.SetOutput(myLog.fileObj)
		//log.SetFlags(log.Ltime | log.Lshortfile)
		//log.Printf(format, v...)
		//fmt.Fprintf(myLog.fileObj, "[%d-%d-%d] [%s:%d] %s", currentTime.Year(), currentTime.Month(), currentTime.Day())
	}

	_, file, line, _ := runtime.Caller(1)
	splitIndex := strings.LastIndex(file, "/")
	fmt.Fprintf(myLog.fileObj, "[%s] [%s:%d] %s\n", currentTime.Local(), file[splitIndex+1:], line, msg)
	fmt.Printf("[%s] [%s:%d] %s\n", currentTime.Local(), file[splitIndex+1:], line, msg)
}

func (myLog *MyLog) Printf(format string, v ...any) {
	var err error
	currentTime := time.Now()
	sFileName := fmt.Sprintf("%d-%d-%d autotest.log", currentTime.Year(), currentTime.Month(), currentTime.Day())
	if sFileName != myLog.fileName {
		myLog.fileName = sFileName
		myLog.fileObj, err = os.OpenFile(myLog.fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			log.Printf("open log file:%s fail", myLog.fileName)
			return
		}
	}
	//caller返回运行信息
	_, file, line, _ := runtime.Caller(1)
	splitIndex := strings.LastIndex(file, "/")
	msg := fmt.Sprintf(format, v...)
	fmt.Fprintf(myLog.fileObj, "[%s] [%s:%d] %s\n", currentTime.Local(), file[splitIndex+1:], line, msg)
	fmt.Printf("[%s] [%s:%d] %s\n", currentTime.Local(), file[splitIndex+1:], line, msg)
}

var Gbs_log MyLog
