package config

import (
	"GoReadNovel/logger"
)

var pythonConfig ConfigInterface

/*
	type GetMaxPageNum struct {
		Page  string `json:"page"`  //页数
		Count string `json:"count"` //数量
	}
	type GetNovelContent struct {
		Title   string `json:"title"`   //页数
		Content string `json:"content"` //数量
	}
*/
// 创建 redis 客户端
func createPythonConfig() ConfigInterface {
	logger.ALogger().Debug("Create Python Config Interface..")
	c, err := NewConfig("./python/python.conf")
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
