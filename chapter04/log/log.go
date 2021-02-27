package log

import (
	"bytes"
	"fmt"
	"log"
)

// 对日志进行设置
func Log() {
	// logger会写入bytes.Buffer类型的数据
	buf := bytes.Buffer{}

	// 第二个参数是前缀最后一个参数是关于选项
	// 配置选项可以用逻辑或符号组合起来
	logger := log.New(&buf, "logger: ", log.Lshortfile|log.Ldate)

	logger.Println("test")

	logger.SetPrefix("new logger: ")

	logger.Printf("you can also add args(%v) and use Fataln to log and crash", true)

	fmt.Println(buf.String())
}
