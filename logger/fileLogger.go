package logger

import (
	"fmt"
	"os"
	"path"
	"time"
)

/*
	打印日志日志的几种形态
	Debug Trace Info Warning Error Fatal
	需求:不同位置输出日志 日志分级别 日志支持开关 日志记录包含时间行号等信息 日志文件要切割
*/
/*
	下面开始写日志类
*/

//file日志类
type FileLogger struct {
	Level       Loglevel
	FilePath    string
	FileName    string
	MaxFileSize int64
	file        *os.File //存放日志文件对象
	errFile     *os.File
}

//日志文件类构造函数
func NewFileLogger(loglevel Loglevel, filePath string, fileName string, maxFileSize int64) (f *FileLogger) {
	if loglevel > FATAL {
		panic("unknow level!")
	}
	f = &FileLogger{
		Level:       loglevel,
		FilePath:    filePath,
		FileName:    fileName,
		MaxFileSize: maxFileSize,
	}
	err := f.initFile()
	if err != nil {
		panic(err)
	}
	return
}

/*
	初始化日志文件

*/
func (f *FileLogger) initFile() (err error) {
	//打开文件
	file, err := os.OpenFile(f.FilePath+f.FileName+".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("open log file failed, err:", err)
		return err
	}
	f.file = file //初始化时就将文件对象存入日志类对象
	errFile, err := os.OpenFile(f.FilePath+f.FileName+"_err.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("open log file failed, err:", err)
		return err
	}
	f.errFile = errFile
	return
}

/*
	关闭日志文件
*/
func (f *FileLogger) Close() {
	f.file.Close()
	f.errFile.Close()
}

/*
	日志打印格式
*/
func (f *FileLogger) log(logLevelName string, format string, a ...interface{}) {
	f.checkFileSizeAndSplitFile()
	str := fmt.Sprintf(format, a...)
	file, funcName, line := getInfo(3)
	now := time.Now().Format("2006-01-02 15:04:05")
	//日志文件
	_, err := fmt.Fprintf(f.file, "[%s] %s [%s %s:%d] %s\n", now, logLevelName, file, funcName, line, str)
	//fmt.Println(n)
	if err != nil {
		fmt.Println("write file log failed , err:", err)
	}

	//错误日志文件
	if logLevelName == "Error" || logLevelName == "Fatal" {
		f.checkErrFileSizeAndSplitErrFile()
		_, err = fmt.Fprintf(f.errFile, "[%s] %s [%s %s:%d] %s\n", now, logLevelName, file, funcName, line, str)
		if err != nil {
			fmt.Println("write file log failed , err:", err)
		}

	}
}

/*
	检查日志文件大小 达到阈值切割文件
	遇到bug 重命名文件失败,因为另一个进程在使用
		可能的原因: fileInfo使用了file 后面调用了fileInfo.Name() 导致找不到
		一定要保证其他方法使用完file对象之后再关闭 比较简单的办法是将关闭旧文件,创建新文件的方法写在一起,中间避免穿插其他对文件进行操作的代码
	这个方法的思想是先关闭旧文件,然后将旧文件重命名,在新建一个新文件对象,将FileLogger结构体的file字段指向新对象
*/
func (f *FileLogger) checkFileSizeAndSplitFile() {
	fileInfo, err := f.file.Stat()
	if err != nil {
		fmt.Println("check fileInfo failed,err:", err)
	}
	//flag := f.MaxFileSize <= fileInfo.Size()
	//fmt.Println(flag)
	if fileInfo.Size() >= f.MaxFileSize {
		// <- 原来关闭文件的方法写在这里
		t := time.Now().Format("_2006_01_02_150405")
		newFileName := f.FilePath + f.FileName + t + ".log"
		oldFileRelName := path.Join(f.FilePath + fileInfo.Name()) //旧文件的全路径名
		f.file.Close()                                            //<- 现在关闭文件的方法,关闭当前文件
		err = os.Rename(oldFileRelName, newFileName)              //修改保存的文件名
		if err != nil {
			fmt.Println("rename log file name failed,err:", err)
		}
		//创建一个新文件
		file, err := os.OpenFile(f.FilePath+f.FileName+".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		//打开文件
		if err != nil {
			fmt.Println("open log file failed, err:", err)
		}
		f.file = file
	}
}

/*
	检查错误日志文件大小 达到阈值切割文件
*/
func (f *FileLogger) checkErrFileSizeAndSplitErrFile() {
	//错误日志大小判断
	errFileInfo, err := f.errFile.Stat()
	if err != nil {
		fmt.Println("check fileInfo failed,err:", err)
	}
	if errFileInfo.Size() >= f.MaxFileSize {

		t := time.Now().Format("_20060102150405")

		newErrFileName := f.FilePath + f.FileName + t + "_err.log"
		//f.FileName = newErrFileName
		oldErrFileRelName := path.Join(f.FilePath + errFileInfo.Name()) //旧文件的全路径名
		//关闭当前文件
		f.errFile.Close()
		err = os.Rename(oldErrFileRelName, newErrFileName) //修改保存的文件名
		if err != nil {
			fmt.Println("rename log file name failed,err:", err)
		}
		errFile, err := os.OpenFile(f.FilePath+f.FileName+"_err.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println("open log file failed, err:", err)
		}
		f.errFile = errFile
	}

}

func (f *FileLogger) Debug(format string, a ...interface{}) {
	if f.Level == DEBUG {
		f.log("Debug", format, a...)
	}
}

func (f *FileLogger) Trace(format string, a ...interface{}) {
	if f.Level <= TRACE {
		f.log("Trace", format, a...)
	}
}
func (f *FileLogger) Info(format string, a ...interface{}) {
	if f.Level <= INFO {
		f.log("Info", format, a...)
	}
}
func (f *FileLogger) Warning(format string, a ...interface{}) {
	if f.Level <= WARNING {
		f.log("Warning", format, a...)
	}
}
func (f *FileLogger) Error(format string, a ...interface{}) {
	if f.Level <= ERROR {
		f.log("Error", format, a...)
	}
}
func (f *FileLogger) Fatal(format string, a ...interface{}) {
	if f.Level <= FATAL {
		f.log("Fatal", format, a...)
	}
}
