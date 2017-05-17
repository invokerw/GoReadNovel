package config

import (
	"GoReadNovel/logger"
)

var pythonConfig ConfigInterface

// 创建 redis 客户端
func createPythonConfig() ConfigInterface {
	logger.ALogger().Debug("Create Python Config Interface..")
	c, err := config.NewConfig("../python/python.conf")
	if err != nil {
		logger.ALogger().Error("create python config error:", err)
		return nil
	}
	return c
}
func init() {
	pythonConfig = createPythonConfig()
}

func GetPythonConfig() ConfigInterface {
	if pythonConfig == nil {
		pythonConfig = createPythonConfig()
	}
	return pythonConfig
}
