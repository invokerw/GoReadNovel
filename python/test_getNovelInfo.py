#!/usr/bin/python
# -*- coding:utf-8 -*-

import urllib2
import sys
import re
import mod_config

reload(sys)
sys.setdefaultencoding('gbk')

# 通过更换strs的不同请求搜索不同的书籍 可以找到对应网页的网址
if len(sys.argv) != 3:
    print 'error : getNovelInfo 参数不够'
    exit()

filename = unicode(sys.argv[1], "UTF-8")
url = unicode(sys.argv[2], "UTF-8")

# strs = u'圣墟'  # 一定要有这个 u  没有u的话请求的编码会有错误
# url = 'http://www.huanyue123.com/book/0/11/'
# values = {'searchkey': strs}
# data = urllib.urlencode(values)

user_agent = 'Mozilla/4.0 (compatible; MSIE 5.5; Windows NT)'# 将user_agent写入头信息
headers = { 'User-Agent' : user_agent }

try:
    request = urllib2.Request(url, headers=headers)
    response = urllib2.urlopen(request,timeout=10)
    content = response.read().decode('gbk')
    head = response.info()
    # print content.encode('utf8')
    mod_config.iniConfig(filename)
    text1 = mod_config.getConfig("getnovelinfo.py", "text1")
	# 描述 类型 最新章节名称 最新章节url 小说状态
    page = re.compile(text1, re.S)
    #page = re.compile('<div.*?class="title">.*?<b>.*?</a>.*?<a.*?>(.*?)</a>.*?<div.*?"options".*?<span.*?'
	#	'<span.*?"item">(.*?)</span>.*?<h3.*?bookinfo_intro.*?</strong>(.*?)<strong>.*?</h3>', re.S)
    replaceBr = re.compile('<br.*?>')
    replaceSpace = re.compile('&nbsp;')
    info = re.findall(page, content.encode('utf8'))

    retStr = ""
    for item in info:
        strTmp = re.sub(replaceBr, "", item[2])
        strTmp = re.sub(replaceSpace, "", strTmp)
        retStr = retStr + item[0] + "-"
        retStr = retStr + item[1] + "-"
        retStr = retStr + strTmp
    print retStr


except urllib2.URLError, e:
    if hasattr(e, "code"):
        print e.code
    if hasattr(e, "reason"):
        print e.reason
