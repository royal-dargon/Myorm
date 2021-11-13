// 这个简易实现的log包，支持日志分级，不同层级的日志显示时使用不同的颜色进行区分，同时显示打印日志代码的文件名和行号
package log

import (
	"io/ioutil"
	"log"
	"os"
	"sync"
)

var (
	// 这里的new是尝试创建一个Logger。参数out设置日志写入的目的地，第二个参数会添加到生成的每一条日志前面，第三个参数定义日志的属性
	// 在这里这个第二个参数确定了颜色，同时使用了log.Lshprtfile 支持显示文件名和代码行号
	errorLog = log.New(os.Stdout, "\033[31m[error]\033[0m ", log.LstdFlags|log.Lshortfile)
	infoLog  = log.New(os.Stdout, "\033[34m[info ]\033[0m ", log.LstdFlags|log.Lshortfile)
	loggers  = []*log.Logger{errorLog, infoLog}
	// 一个互斥锁
	mu sync.Mutex
)

// log methods
var (
	Error  = errorLog.Println
	Errorf = errorLog.Printf
	Info   = infoLog.Println
	Infof  = infoLog.Printf
)

// 下面开始支持设置日志的层数
// log levels
const (
	InfoLevel = iota
	ErrorLevel
	Disabled
)

// SetLevel controls log level
func SetLevel(level int) {
	mu.Lock()
	defer mu.Unlock()

	for _, logger := range loggers {
		// 设置标准logger的输出目的地
		logger.SetOutput(os.Stdout)
	}

	if ErrorLevel < level {
		// 不做操作
		infoLog.SetOutput(ioutil.Discard)
	}
	if InfoLevel < level {
		infoLog.SetOutput(ioutil.Discard)
	}
}

// 以上部分就是决定了三个层级是否进行打印
// 然后log就完成了
