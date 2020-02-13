package main

import (
	"fmt"
	"github.com/wulusai2333/gostudy/logger"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
	"time"
)

/*
	time包
	日期的各种表示方法
	日期运算 日期格式化 日期转换
*/
func main() {
	//时间包使用
	//time_test()
	//控制台日志
	//log := logger.NewLogger(logger.INFO)
	//文件日志
	log := logger.NewFileLogger(logger.INFO, "logger/", "test", 1*1024)
	fmt.Println(log)
	//GetAppPath() //路径问题
	for {
		log.Debug("this is a Debug log")
		log.Trace("this is a Trace log")
		log.Info("this is a Info log")
		log.Warning("this is a Warning log")
		log.Error("this is a Error log")
		log.Fatal("this is a Fatal log")
		//time.Sleep(time.Second)
	}
}

func time_test() {
	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())
	//时间戳 秒 纳秒
	fmt.Println(now.Unix())
	fmt.Println(now.UnixNano())
	//时间戳转换
	t := time.Unix(now.Unix(), 1000)
	fmt.Println(t)
	fmt.Println(time.Second)
	//go语言的时间格式化与众不同 2006 1 2 3 4 5 代表其他语言的YY-MM-DD HH:mm:ss
	fmt.Println(now.Format("2006-01-02 03:04:05"))
	fmt.Println(now.Format("2006-01-02 15:04:05"))
	fmt.Println(now.Format("2006-01-02 15:04:05 PM"))
	fmt.Println(now.Format("2006-01-02 15:04:05.000 PM"))
	//指定格式字符串转成时间
	t, err := time.Parse("2006/01/02 03 04 05", "2020/05/01 11 17 59")
	if err != nil {
		fmt.Println("parse time failed, err:", err)
		return
	}
	fmt.Println(t)
	//时区格式化
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println("load location failed, err:", err)
		return
	}
	//按照时区解析时间
	t, err = time.ParseInLocation("2006/01/02 03 04 05", "2020/05/01 11 17 59", loc)
	if err != nil {
		fmt.Println("parse time in location failed,err:", err)
		return
	}
	fmt.Println(t)
	//时间相减
	fmt.Println(t.Sub(now))
}
func GetAppPath() {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	index := strings.LastIndex(path, string(os.PathSeparator))
	fmt.Println("当前可执行程序路径:", path[:index])
	path, _ = os.Getwd()
	fmt.Println("当前用户路径:", path)
	path, _ = exec.LookPath(os.Args[0])
	fmt.Println("执行程序文件相对路径:", path)
	relPath()
	return
}

func relPath() {
	file, _ := exec.LookPath(os.Args[0])
	log.Println("file:", file)
	dir, _ := path.Split(file)
	log.Println("dir:", dir)
	os.Chdir(dir)
	wd, _ := os.Getwd()
	log.Println("wd:", wd)
}
