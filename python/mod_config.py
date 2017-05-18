#encoding:utf-8
#name:mod_config.py

import ConfigParser
import os

#获取config配置文件

CONFIG = None

def iniConfig(filename):
    global CONFIG
    CONFIG = ConfigParser.ConfigParser()
    path = os.path.split(os.path.realpath(__file__))[0] + '/'+ filename
    CONFIG.read(path)
def getConfig(section, key):
    return CONFIG.get(section, key)

#其中 os.path.split(os.path.realpath(__file__))[0] 得到的是当前文件模块的目录
