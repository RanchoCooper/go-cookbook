package log

import "github.com/pkg/errors"
import "log"

// OriginalError返回错误的原始信息
func OriginalError() error {
	return errors.New("error occurred")
}

// PassThroughError 调用OriginalError并将其封装
func PassThroughError() error {
	err := OriginalError()
	// 无需检查错误，因为使用该库时此操作适用于nil
	return errors.Wrap(err, "in passthrougherror")
}

// FinalDestination处理错误而不传递它
func FinalDestination() {
	err := PassThroughError()
	if err != nil {
		// 将任何产生的意外记录到日志中
		log.Printf("an error occurred: %s\n", err.Error())
		return
	}
}
