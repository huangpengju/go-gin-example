package logging

import (
	"fmt"
	"log"
	"os"
	"time"
)

var (
	LogSavePath = "runtime/logs/" // log保存路径
	LogSaveName = "log"           // log文件名称
	LogFileExt  = "log"           // log文件的后缀名
	TimeFormat  = "20060102"      // 时间格式
)

// 获取 log 文件路径
func getLogFilePath() string {
	// Sprintf根据格式说明符格式化并返回结果字符串。
	// return fmt.Sprintf("%s", LogSavePath)
	return LogSavePath
}

// 获取 log 文件完整路径（包括文件名）
func getLogFileFullPath() string {
	// 获取 log 文件路径
	prefixPath := getLogFilePath()
	// 拼接生成 log 文件全名
	suffixPath := fmt.Sprintf("%s%s.%s", LogSaveName, time.Now().Format(TimeFormat), LogFileExt)
	// 返回 log 文件路径及文件名
	return fmt.Sprintf("%s%s", prefixPath, suffixPath)
}

// openLogFile 获取文件指针
// filePath 是文件的完整路径（包含文件名）
func openLogFile(filePath string) *os.File {
	// 返回文件信息结构描述文件。如果出现错误，会返回*PathError
	_, err := os.Stat(filePath)
	switch {
	// os.IsNotExist：能够接受ErrNotExist、syscall的一些错误，它会返回一个布尔值，能够得知文件不存在或目录不存在
	case os.IsNotExist(err):
		mkDir()
	// os.IsPermission：能够接受ErrPermission、syscall的一些错误，它会返回一个布尔值，能够得知权限是否满足
	case os.IsPermission(err):
		log.Fatalf("Permission:%v", err)
	}
	// os.OpenFile：调用文件，支持传入文件名称、指定的模式调用文件、文件权限，返回的文件的方法可以用于 I/O。如果出现错误，则为*PathError。
	// O_APPEND 在写入时将数据追加到文件中
	// O_CREATE 如果不存在，则创建一个新文件
	// O_WRONLY 以只写模式打开文件
	handle, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("打开文件：%v 失败", err)
	}
	return handle
}

// mkDir 创建文件夹以及子文件
func mkDir() {
	// os.Getwd：返回与当前目录对应的根路径名
	dir, _ := os.Getwd()
	// os.MkdirAll：创建对应的目录以及所需的子目录，若成功则返回nil，否则返回error
	// getLogFilePath() 获取 log 文件路径
	// os.ModePerm：const定义ModePerm FileMode = 0777
	err := os.MkdirAll(dir+"/"+getLogFilePath(), os.ModePerm)
	if err != nil {
		panic(err)
	}
}
