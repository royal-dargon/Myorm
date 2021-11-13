// 这个简易实现的log包，支持日志分级，不同层级的日志显示时使用不同的颜色进行区分，同时显示打印日志代码的文件名和行号
package log

import (
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
