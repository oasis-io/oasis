package logs

import (
	"go.uber.org/zap"
)

var logs *zap.Logger

func init() {
	var err error
	logs, err = zap.NewProduction()
	if err != nil {
		// 记录错误日志等
		logs = zap.NewExample()
	}
}

func Logger() *zap.Logger {
	if logs == nil {
		// 如果未初始化，可以尝试重新初始化或返回默认 Logger 对象等
		return zap.NewExample()
	}
	return logs
}
