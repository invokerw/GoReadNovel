package config

import (
	"GoReadNovel/logger"
	"strconv"
)

var pythonConfig ConfigInterface

//参考python.conf
type PythonConfig struct {
	Name           string   `json:"name"`           //名称
	Info           string   `json:"info"`           //数量
	TestValueCount int      `json:"testvaluecount"` //需要测试变量数量
	TestValues     []string `json:"testvalues"`     //测试变量
	TextCount      int      `json:"textcount"`      //正则表达式数量
	TextValues     []string `json:"textvalues"`     //正则表达式内容
}

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

func GetPythonConfigInterface() ConfigInterface {
	if pythonConfig == nil {
		pythonConfig = createPythonConfig()
	}
	return pythonConfig
}
func GetAPythonPageConfig(id int) (PythonConfig, bool) {
	conf := PythonConfig{}
	pagecount, err := GetPythonConfigInterface().Int("pagecount")
	if err != nil {
		logger.ALogger().Error("GetAPythonPageConfig get pagecount error:", err)
		return conf, false
	}
	if pagecount < id {
		return conf, false
	}
	idStr := strconv.Itoa(id)
	tmpStr := idStr + "::name"
	conf.Name = GetPythonConfigInterface().String(tmpStr)
	tmpStr = idStr + "::info"
	conf.Info = GetPythonConfigInterface().String(tmpStr)
	tmpStr = idStr + "::testvaluecount"
	conf.TestValueCount, err = GetPythonConfigInterface().Int(tmpStr)
	if err != nil {
		logger.ALogger().Error("GetAPythonPageConfig get testvaluecount error:", err)
		return conf, false
	}

	for i := 1; i <= conf.TestValueCount; i++ {
		tmpStr = idStr + "::testvalue" + strconv.Itoa(i)
		tmpvalue := ""
		tmpvalue = GetPythonConfigInterface().String(tmpStr)
		conf.TestValues = append(conf.TestValues, tmpvalue)
	}

	tmpStr = conf.Name + "::textcount"
	conf.TextCount, err = GetPythonConfigInterface().Int(tmpStr)
	if err != nil {
		logger.ALogger().Error("GetAPythonPageConfig get textcount error:", err)
		return conf, false
	}
	for i := 1; i <= conf.TextCount; i++ {
		tmpStr = conf.Name + "::text" + strconv.Itoa(i)
		textvalue := ""
		textvalue = GetPythonConfigInterface().String(tmpStr)
		conf.TextValues = append(conf.TextValues, textvalue)
	}
	return conf, true
}
