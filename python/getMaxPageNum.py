#!/usr/bin/python
# -*- coding:utf-8 -*-
import urllib
import urllib2
import sys
import re
#获取小说的数量及页数
reload(sys)
sys.setdefaultencoding('gbk')

url = 'http://www.huanyue123.com/book/'

user_agent = 'Mozilla/4.0 (compatible; MSIE 5.5; Windows NT)'# 将user_agent写入头信息

headers = { 'User-Agent' : user_agent }


try:
    request = urllib2.Request(url,headers=headers)
    response = urllib2.urlopen(request,timeout=5)
    content = response.read().decode('gbk')
    head = response.info()
    # print content

    page = re.compile('<em.*?id="pagestats.*?1/(.*?)</em>"', re.S)
    novel = re.compile('<h1.*?class="title">.*?共有小说(.*?)本.*?</b>.*?</h1>', re.S)

    pageNum = re.findall(page, content.encode('utf8'))
    novelNum = re.findall(novel, content.encode('utf8'))

    print pageNum,novelNum

    retStr = ""

    print retStr


except urllib2.URLError, e:
    if hasattr(e, "code"):
        print e.code
    if hasattr(e, "reason"):
        print e.reason
