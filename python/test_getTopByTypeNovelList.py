#!/usr/bin/python
# -*- coding:utf-8 -*-

import urllib2
import sys
import re
import mod_config

reload(sys)
sys.setdefaultencoding('gbk')


if len(sys.argv) != 5:
    print 'Err : getTopByTypeNovelList参数不够'
    exit()

filename = unicode(sys.argv[1], "UTF-8")
novelType = unicode(sys.argv[2], "UTF-8")
sortType = unicode(sys.argv[3], "UTF-8")
page = unicode(sys.argv[4], "UTF-8")

url = 'http://www.huanyue123.com/book/' + novelType + '/' + sortType + '-0-0-0-0-0-0-' + page + '.html'
# url = 'http://www.huanyue123.com/book/quanbu/allvisit-0-0-0-0-0-0-1.html'
# print url

user_agent = 'Mozilla/4.0 (compatible; MSIE 5.5; Windows NT)'# 将user_agent写入头信息
headers = { 'User-Agent' : user_agent }
try:
    request = urllib2.Request(url,headers=headers)
    response = urllib2.urlopen(request,timeout=5)
    content = response.read().decode('gbk')
    head = response.info()
    # print content
    mod_config.iniConfig(filename)
    text1 = mod_config.getConfig("getTopByTypeNovelList.py", "text1")
    page = re.compile(text1, re.S)
    # page = re.compile(
    #    '<dl>.*?<dt>.*?<img.*?src="(.*?)".*?>.*?</dt>.*?<dd>.*?<h3>.*?<span.*?"uptime">(.*?)</span>.*?'
    #    '<a.*?href="(.*?)".*?>(.*?)</a>.*?</h3>.*?</dd>.*?<span>(.*?)</span>.*?'
    #    '<dd.*?>(.*?)</dd>.*?<dd.*?<a.*?href="(.*?)".*?>(.*?)</a>.*?</dd>.*?'
    #    '</dl>', re.S)
    # item[2] 小说地址, item[3] 小说名字, item[6]最新章节地址(可用),
    # item[7]最新章节名字, item[4] 作者, item[1]更新时间, item[5]描述
    # item[0]描述
    hrefList = re.findall(page, content.encode('utf8'))

    chapterQty = 0
    # print "Begin->"
    retStr = ""
    for item in hrefList:
        chapterQty = chapterQty + 1
        retStr = retStr + str(chapterQty) + "--" + item[1]
        retStr = retStr + "--" + item[2] + "--" + item[3]
        retStr = retStr + "--" + item[4] + "--" + item[5]
        retStr = retStr + "--" + item[6] + "--" + item[7]
        retStr = retStr + "--" + item[0] + ","
        #print item[0], item[1], item[2], item[3], item[4], item[5], item[6]
    print retStr


except urllib2.URLError, e:
    if hasattr(e, "code"):
        print e.code
    if hasattr(e, "reason"):
        print e.reason
