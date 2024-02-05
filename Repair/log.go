package repair

import (
	"fmt"
	"os"
)

// 终端颜色
var (
	DEBUG   string = "[\033[1;34mD\033[0m]"
	INFO    string = "[\033[1;35m*\033[0m]"
	WARNING string = "[\033[0;38;5;214m!\033[0m]"
	ERR     string = "[\033[1;31m✘\033[0m]"
	RIGHT   string = "[\033[1;32m✓\033[0m]"
	GREY    string = "\033[1;30m"
	RED     string = "\033[1;31m"
	GREEN   string = "\033[1;32m"
	YELLOW  string = "\033[1;33m"
	BLUE    string = "\033[1;34m"
	PURPLE  string = "\033[1;35m"
	MAIN    string = "\033[1;36m"
	ORANGE  string = "\033[0;38;5;214m"
	END     string = "\033[0m"
)

// 写入文件
func writeFile(log string, filename string) {
	var text = []byte(log + "\n")
	fl, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("%s 打开 %s 失败，%v\n", ERR, filename, err)
		return
	}
	defer fl.Close()
	_, err = fl.Write(text)
	if err != nil {
		fmt.Printf("%s 写入 %s 失败，%v\n", ERR, filename, err)
	}
}

// 打印内容
func PrintLog(format string, args ...interface{}) {
	logdata := fmt.Sprintf(format, args...)
	writeFile("    "+logdata, outfile)
	fmt.Printf("    %s\n", logdata)
}

// 设置必须显示的日志颜色
func InfoLog(format string, args ...interface{}) {
	logdata := fmt.Sprintf(format, args...)
	writeFile("[*] "+logdata, outfile)
	fmt.Printf("%s %s\n", INFO, logdata)
}

// 设置正确颜色
func RightLog(format string, args ...interface{}) {
	logdata := fmt.Sprintf(format, args...)
	writeFile("[✓] "+logdata, outfile)
	fmt.Printf("%s %s\n", RIGHT, logdata)
}

// 设置必须的错误颜色
func ErrorLog(format string, args ...interface{}) {
	logdata := fmt.Sprintf(format, args...)
	writeFile("[✘] "+logdata, outfile)
	fmt.Printf("%s %s\n", ERR, logdata)
}

// 设置错误颜色
func ErrLog(format string, args ...interface{}) {
	if debugLevel >= 1 {
		logdata := fmt.Sprintf(format, args...)
		writeFile("[✘] "+logdata, outfile)
		fmt.Printf("%s %s\n", ERR, logdata)
	}
}

// 设置debug等级
func DebugLog(format string, args ...interface{}) {
	if debugLevel >= 2 {
		logdata := fmt.Sprintf(format, args...)
		writeFile("[D] "+logdata, outfile)
		fmt.Printf("%s %s\n", DEBUG, logdata)
	}
}

// 设置警告
func WarningLog(format string, args ...interface{}) {
	if debugLevel >= 3 {
		logdata := fmt.Sprintf(format, args...)
		writeFile("[!] "+logdata, outfile)
		fmt.Printf("%s %s\n", WARNING, logdata)
	}
}
