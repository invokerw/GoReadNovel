#!/usr/bin/python
# -*- coding:utf-8 -*-
import urllib
import urllib2
import sys
import re

reload(sys)
sys.setdefaultencoding('gbk')

# 通过更换strs的不同请求搜索不同的书籍 可以找到对应网页的网址
if len(sys.argv) != 2:
    print 'Err : GetNoteTxt参数不够'
    exit()
strs = unicode(sys.argv[1], "UTF-8")
# strs = u'圣墟'  # 一定要有这个 u  没有u的话请求的编码会有错误

url = 'http://www.huanyue123.com/modules/article/search.php'

values = {'searchkey': strs}

data = urllib.urlencode(values)


try:
    request = urllib2.Request(url, data=data)
    response = urllib2.urlopen(request)
    content = response.read().decode('gbk')
    head = response.info()

    # page = re.compile('<li>.*?<a.*?href="(.*?)">(.*?)</a>*?</li>', re.S)
    page = re.compile(
        '<tr>.*?<td.*?href="(.*?)">(.*?)<.*?href="(.*?)".*?>'
        '(.*?)</a.*?<td.*?>(.*?)</.*?<td.*?</td>.*?<td.*?>(.*?)</td>'
        '.*?<td.*?>(.*?)</td>', re.S)
    # item[0] 小说地址, item[1] 小说名字, item[2]最新章节地址(不可用),
    # item[3]最新章节名字, item[4] 作者, item[5]跟新时间, item[6]状态 连载还是完结
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
        # print item[0], item[1], item[2], item[3], item[4], item[5]
    print retStr


except urllib2.URLError, e:
    if hasattr(e, "code"):
        print e.code
    if hasattr(e, "reason"):
        print e.reason
