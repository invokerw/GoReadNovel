#!/usr/bin/python
# -*- coding:utf-8 -*-

import urllib2
import sys
import re

reload(sys)
sys.setdefaultencoding('gbk')

url = 'http://www.huanyue123.com/book/top.html'
user_agent = 'Mozilla/4.0 (compatible; MSIE 5.5; Windows NT)'# 将user_agent写入头信息
headers = { 'User-Agent' : user_agent }
try:
    request = urllib2.Request(url,headers=headers)
    response = urllib2.urlopen(request)
    content = response.read().decode('gbk')
    head = response.info()

    # page = re.compile('<li>.*?<a.*?href="(.*?)">(.*?)</a>*?</li>', re.S)
    page = re.compile(
        '<li>.*?<p.*?class="s1".*?<a.*?href="(.*?)".*?>(.*?)</a>.*?<p.*?<a.*?'
        'href="(.*?)".*?>(.*?)</a>.*?<p.*?>(.*?)</p>.*?<p.*?/p>.*?<p.*?>(.*?)'
        '</p>.*?<p.*?>(.*?)</p></li>', re.S)
    # item[0] 小说地址, item[1] 小说名字, item[2]最新章节地址(可用),
    # item[3]最新章节名字, item[4] 作者, item[6]更新时间, item[5]状态 连载还是完结
    hrefList = re.findall(page, content.encode('utf8'))

    chapterQty = 0
    # print "Begin->"
    retStr = ""
    for item in hrefList:
        chapterQty = chapterQty + 1
        retStr = retStr + str(chapterQty) + "--" + item[0]
        retStr = retStr + "--" + item[1] + "--" + item[3]
        retStr = retStr + "--" + item[4] + "--" + item[5]
        retStr = retStr + "--" + item[2] + ","
        # print item[0], item[1], item[2], item[3], item[4], item[5], item[6]
    print retStr


except urllib2.URLError, e:
    if hasattr(e, "code"):
        print e.code
    if hasattr(e, "reason"):
        print e.reason
