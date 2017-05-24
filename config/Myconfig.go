package config

import (
	"GoReadNovel/logger"
	"strconv"
	"strings"
)

/*
var pythonConfig ConfigInterface
var testConfig ConfigInterface
*/
var pyConfFile *ConfigFile
var testConfFile *ConfigFile
var err error

//参考python.conf
type PythonConfig struct {
	ID             int      `json:"id"`             //ID
	Name           string   `json:"name"`           //名称
	Info           string   `json:"info"`           //数量
	TestValueCount int      `json:"testvaluecount"` //需要测试变量数量
	TestValues     []string `json:"testvalues"`     //测试变量
	TextCount      int      `json:"textcount"`      //正则表达式数量
	TextValues     []string `json:"textvalues"`     //正则表达式内容
}

/*
func createPythonConfig() ConfigInterface {
	logger.ALogger().Debug("Create Python Config Interface..")
	c, err := NewConfig("./python/python.conf")
	if err != nil {
		logger.ALogger().Error("create python config error:", err)
		return nil
	}
	return c
}
*/
func createPythonConfig() {
	logger.ALogger().Debug("Create Python Config File..")
	testConfFile, err = LoadConfigFile("./python/python.conf")
	if err != nil {
		logger.ALogger().Error("create python config error:", err)
	}
}

/*
func createTestConfig() ConfigInterface {
	logger.ALogger().Debug("Create Test Python Config Interface..")
	c, err := NewConfig("./python/test.conf")
	if err != nil {
		logger.ALogger().Error("create test python config error:", err)
		return nil
	}
	return c
}
*/
func createTestConfig() {
	logger.ALogger().Debug("Create Test Python Config File..")
	pyConfFile, err = LoadConfigFile("./python/test.conf")
	if err != nil {
		logger.ALogger().Error("create test python config error:", err)
	}
}
func init() {
	createPythonConfig()
	createTestConfig()
}
func TestConfigReload() {
	err := testConfFile.Reload()
	if err != nil {
		logger.ALogger().Error("testConfFile reload error:", err)
	}
}

/*
func GetTestConfigInterface() ConfigInterface {
	if testConfig == nil {
		testConfig = createTestConfig()
	}
	return testConfig
}
*/
func GetTestConfigInterface() *ConfigFile {
	if testConfFile == nil {
		createTestConfig()
	}
	return testConfFile
}

/*
func GetPythonConfigInterface() ConfigInterface {
	if pythonConfig == nil {
		pythonConfig = createPythonConfig()
	}
	return pythonConfig
}
*/
func GetPythonConfigInterface() *ConfigFile {
	if pyConfFile == nil {
		createPythonConfig()
	}
	return pyConfFile
}

/*
func GetAPythonPageConfig(id int) (PythonConfig, bool) {
	conf := PythonConfig{}
	pagecount, err := GetPythonConfigInterface().Int("pagecount::pagecount")
	if err != nil {
		logger.ALogger().Error("GetAPythonPageConfig get pagecount error:", err)
		return conf, false
	}
	if pagecount < id {
		return conf, false
	}
	conf.ID = id
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
		//return conf, false
	}
	for i := 1; i <= conf.TextCount; i++ {
		tmpStr = conf.Name + "::text" + strconv.Itoa(i)
		textvalue := ""
		textvalue = GetPythonConfigInterface().String(tmpStr)
		conf.TextValues = append(conf.TextValues, textvalue)
	}
	return conf, true
}
*/
//重写
func GetAPythonPageConfig(id int) (PythonConfig, bool) {
	conf := PythonConfig{}
	pagecount, err := GetPythonConfigInterface().Int("pagecount", "pagecount")
	if err != nil {
		logger.ALogger().Error("GetAPythonPageConfig get pagecount error:", err)
		return conf, false
	}
	if pagecount < id {
		return conf, false
	}
	conf.ID = id
	idStr := strconv.Itoa(id)

	conf.Name, _ = GetPythonConfigInterface().GetValue(idStr, "name")
	conf.Info, _ = GetPythonConfigInterface().GetValue(idStr, "info")
	conf.TestValueCount, err = GetPythonConfigInterface().Int(idStr, "testvaluecount")
	if err != nil {
		logger.ALogger().Error("GetAPythonPageConfig get testvaluecount error:", err)
		return conf, false
	}
	tmpStr := ""
	for i := 1; i <= conf.TestValueCount; i++ {
		tmpStr = "testvalue" + strconv.Itoa(i)
		tmpvalue := ""
		tmpvalue, _ = GetPythonConfigInterface().GetValue(idStr, tmpStr)
		conf.TestValues = append(conf.TestValues, tmpvalue)
	}
	title := strings.ToLower(conf.Name)
	conf.TextCount, err = GetPythonConfigInterface().Int(title, "textcount")
	if err != nil {
		logger.ALogger().Error("GetAPythonPageConfig get textcount error:", err)
		//return conf, false
	}
	for i := 1; i <= conf.TextCount; i++ {
		tmpStr = "text" + strconv.Itoa(i)
		textvalue := ""
		textvalue, _ = GetPythonConfigInterface().GetValue(title, tmpStr)
		conf.TextValues = append(conf.TextValues, textvalue)
	}
	return conf, true
}

/*
func WriteAPythonPageConfig(conf PythonConfig) bool {

	pagecount, err := GetPythonConfigInterface().Int("pagecount::pagecount")
	if err != nil {
		logger.ALogger().Error("WriteAPythonPageConfig get pagecount error:", err)
		return false
	}
	if pagecount < conf.ID {
		return false
	}
	idStr := strconv.Itoa(conf.ID)
	tmpStr := idStr + "::name"
	GetPythonConfigInterface().Set(tmpStr, conf.Name)
	tmpStr = idStr + "::info"
	GetPythonConfigInterface().Set(tmpStr, conf.Info)
	tmpStr = idStr + "::testvaluecount"
	GetPythonConfigInterface().Set(tmpStr, strconv.Itoa(conf.TestValueCount))

	for i := 1; i <= conf.TestValueCount; i++ {
		tmpStr = idStr + "::testvalue" + strconv.Itoa(i)
		GetPythonConfigInterface().Set(tmpStr, conf.TestValues[i-1])
	}

	tmpStr = conf.Name + "::textcount"
	GetPythonConfigInterface().Set(tmpStr, strconv.Itoa(conf.TextCount))

	for i := 1; i <= conf.TextCount; i++ {
		tmpStr = conf.Name + "::text" + strconv.Itoa(i)
		GetPythonConfigInterface().Set(tmpStr, conf.TextValues[i-1])
	}
	return true
}
func WriteATestPageConfig(conf PythonConfig) bool {

	pagecount, err := GetTestConfigInterface().Int("pagecount::pagecount")
	if err != nil {
		logger.ALogger().Error("WriteATestPageConfig get pagecount error:", err)
		return false
	}
	if pagecount < conf.ID {
		return false
	}
	idStr := strconv.Itoa(conf.ID)
	tmpStr := idStr + "::name"
	logger.ALogger().Debugf("%s = %s", tmpStr, conf.Name+"xxx")
	err = GetTestConfigInterface().Set(tmpStr, conf.Name+"xxx")
	if err != nil {
		logger.ALogger().Error("WriteATestPageConfig Set name error:", err)
		return false
	}
	tmpStr = idStr + "::info"
	GetTestConfigInterface().Set(tmpStr, conf.Info)

	tmpStr = idStr + "::testvaluecount"
	GetTestConfigInterface().Set(tmpStr, strconv.Itoa(conf.TestValueCount))

	for i := 1; i <= conf.TestValueCount; i++ {
		tmpStr = idStr + "::testvalue" + strconv.Itoa(i)
		GetTestConfigInterface().Set(tmpStr, conf.TestValues[i-1])
	}

	tmpStr = conf.Name + "::textcount"
	GetTestConfigInterface().Set(tmpStr, strconv.Itoa(conf.TextCount))
	if err != nil {
		logger.ALogger().Error("set key err", err)
		return false
	}

	for i := 1; i <= conf.TextCount; i++ {
		tmpStr = conf.Name + "::text" + strconv.Itoa(i)
		GetTestConfigInterface().Set(tmpStr, conf.TextValues[i-1])
	}
	return true
}
*/
func WriteAPythonPageConfig(conf PythonConfig) bool {

	pagecount, err := GetPythonConfigInterface().Int("pagecount", "pagecount")
	if err != nil {
		logger.ALogger().Error("WriteAPythonPageConfig get pagecount error:", err)
		return false
	}
	if pagecount < conf.ID {
		return false
	}
	idStr := strconv.Itoa(conf.ID)
	GetPythonConfigInterface().SetValue(idStr, "name", conf.Name)
	GetPythonConfigInterface().SetValue(idStr, "info", conf.Info)
	GetPythonConfigInterface().SetValue(idStr, "testvaluecount", strconv.Itoa(conf.TestValueCount))
	tmpStr := ""
	for i := 1; i <= conf.TestValueCount; i++ {
		tmpStr = "testvalue" + strconv.Itoa(i)
		GetPythonConfigInterface().SetValue(idStr, tmpStr, conf.TestValues[i-1])
	}
	title := strings.ToLower(conf.Name)
	GetPythonConfigInterface().SetValue(title, "textcount", strconv.Itoa(conf.TextCount))

	for i := 1; i <= conf.TextCount; i++ {
		tmpStr = "text" + strconv.Itoa(i)
		GetPythonConfigInterface().SetValue(title, tmpStr, conf.TextValues[i-1])
	}
	return true
}
func WriteATestPageConfig(conf PythonConfig) bool {

	pagecount, err := GetTestConfigInterface().Int("pagecount", "pagecount")
	if err != nil {
		logger.ALogger().Error("WriteAPythonPageConfig get pagecount error:", err)
		return false
	}
	if pagecount < conf.ID {
		return false
	}
	idStr := strconv.Itoa(conf.ID)
	GetTestConfigInterface().SetValue(idStr, "name", conf.Name)
	GetTestConfigInterface().SetValue(idStr, "info", conf.Info)
	GetTestConfigInterface().SetValue(idStr, "testvaluecount", strconv.Itoa(conf.TestValueCount))
	tmpStr := ""
	for i := 1; i <= conf.TestValueCount; i++ {
		tmpStr = "testvalue" + strconv.Itoa(i)
		GetTestConfigInterface().SetValue(idStr, tmpStr, conf.TestValues[i-1])
	}
	title := strings.ToLower(conf.Name)
	GetTestConfigInterface().SetValue(title, "textcount", strconv.Itoa(conf.TextCount))

	for i := 1; i <= conf.TextCount; i++ {
		tmpStr = "text" + strconv.Itoa(i)
		GetTestConfigInterface().SetValue(title, tmpStr, conf.TextValues[i-1])
	}
	err = SaveConfigFile(GetTestConfigInterface(), "./python/test.conf")
	return true
}
