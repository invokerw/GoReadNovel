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
    print 'error : getNovelChapterByUrl 参数不够'
    exit()
filename = unicode(sys.argv[1], "UTF-8")
url = unicode(sys.argv[2], "UTF-8")

user_agent = 'Mozilla/4.0 (compatible; MSIE 5.5; Windows NT)'# 将user_agent写入头信息
headers = { 'User-Agent' : user_agent }

try:
    request = urllib2.Request(url, headers=headers)
    response = urllib2.urlopen(request)
    content = response.read().decode('gbk')
    head = response.info()
    
    mod_config.iniConfig(filename)
    text1 = mod_config.getConfig("getnovelchapterbyurl.py", "text1")
    page = re.compile(text1, re.S)
    # page = re.compile('<li>.*?<a.*?href="(.*?)">(.*?)</a>*?</li>', re.S)

    hrefList = re.findall(page, content.encode('utf8'))
    chapterQty = 0

    retStr = ""
    for item in hrefList:
        chapterQty = chapterQty + 1
        retStr = retStr + str(chapterQty) + "-" + item[0]
        retStr = retStr + "-" + item[1] + ","

    print retStr


except urllib2.URLError, e:
    if hasattr(e, "code"):
        print e.code
    if hasattr(e, "reason"):
        print e.reason
