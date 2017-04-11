#!/usr/bin/python
# -*- coding:utf-8 -*-

import urllib2
import sys
import re

reload(sys)
sys.setdefaultencoding('gbk')


if len(sys.argv) != 4:
    print 'Err : GetNovelTxt参数不够'
    exit()
novelType = unicode(sys.argv[1], "UTF-8")
sortType = unicode(sys.argv[2], "UTF-8")
page = unicode(sys.argv[3], "UTF-8")

url = 'http://www.huanyue123.com/book/' + novelType + '/' + sortType + '-0-0-0-0-0-0-' + page + '.html'
# url = 'http://www.huanyue123.com/book/quanbu/allvisit-0-0-0-0-0-0-1.html'

# print url

user_agent = 'Mozilla/4.0 (compatible; MSIE 5.5; Windows NT)'# 将user_agent写入头信息
headers = { 'User-Agent' : user_agent }
try:
    request = urllib2.Request(url,headers=headers)
    response = urllib2.urlopen(request)
    content = response.read().decode('gbk')
    head = response.info()
    # print content

    # page = re.compile('<li>.*?<a.*?href="(.*?)">(.*?)</a>*?</li>', re.S)
    page = re.compile(
        '<dl>.*?<dt>.*?</dt>.*?<dd>.*?<h3>.*?<span.*?"uptime">(.*?)</span>.*?'
        '<a.*?href="(.*?)".*?>(.*?)</a>.*?</h3>.*?</dd>.*?<span>(.*?)</span>.*?'
        '<dd.*?>(.*?)</dd>.*?<dd.*?<a.*?href="(.*?)".*?>(.*?)</a>.*?</dd>.*?'
        '</dl>', re.S)
    # item[1] 小说地址, item[2] 小说名字, item[5]最新章节地址(可用),
    # item[6]最新章节名字, item[3] 作者, item[0]更新时间, item[4]描述
    hrefList = re.findall(page, content.encode('utf8'))

    chapterQty = 0
    # print "Begin->"
    retStr = ""
    for item in hrefList:
        chapterQty = chapterQty + 1
        retStr = retStr + str(chapterQty) + "--" + item[0]
        retStr = retStr + "--" + item[1] + "--" + item[2]
        retStr = retStr + "--" + item[3] + "--" + item[4]
        retStr = retStr + "--" + item[5] + "--" + item[6]
        retStr = retStr +","
        #print item[0], item[1], item[2], item[3], item[4], item[5], item[6]
    print retStr


except urllib2.URLError, e:
    if hasattr(e, "code"):
        print e.code
    if hasattr(e, "reason"):
        print e.reason
