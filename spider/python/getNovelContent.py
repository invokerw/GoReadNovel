#!/usr/bin/python
# -*- coding:utf-8 -*-
import urllib2
import sys
import re

reload(sys)
sys.setdefaultencoding('gbk')

# 输入IP地址转换为文章文件保存至当前目录下  后面可能需要改一下输入保存地址

if len(sys.argv) != 2:
    print 'Err : GetNovelContent参数不够'
    exit()
url = unicode(sys.argv[1], "UTF-8")
# url = 'http://www.huanyue123.com/book/0/11/925296.html'  # 文章的地址
# print url
user_agent = 'Mozilla/4.0 (compatible; MSIE 5.5; Windows NT)'# 将user_agent写入头信息
headers = { 'User-Agent' : user_agent }
try:
    request = urllib2.Request(url,headers=headers)
    response = urllib2.urlopen(request)
    content = response.read().decode('gbk')
    # head = response.info()
    title = re.compile('<div.*?class="h1title.*?>.*?<h1.*?>(.*?)</h1>' +
                       '.*?</div>', re.S)
    article = re.compile('<div.*?id="htmlContent".*?>(.*?)</div>', re.S)

    urlpre = re.compile('<a.*?href="(.*?)".*?>上一章.*?章节目录.*?加入书签.*?投票推荐.*?'
                        '<a.*?href="(.*?)".*?>下一章</a>', re.S)
    nName = re.compile('<div.*?"title".*?>.*?<strong>(.*?)</strong>', re.S)

    # 将<br />替换为\t  这里不做替换 因为是要在网页显示
    replaceBr = re.compile('<br.*?\n.*?<br.*?>')
    replaceSpace = re.compile('&nbsp;')

    replaceBrIndex = re.compile(u'章节目录 ')

    tit = ''    # 文章名字
    arti = ''   # 文章内容
    uPre = ''   # 上一章url
    uNext = ''  # 下一章url
    uName = ''  # 小说名字

    itemsTit = re.findall(title, content.encode('utf8'))
    for item in itemsTit:
        tit = item

    itemsArticle = re.findall(article, content.encode('utf8'))
    for item in itemsArticle:
        arti = item

    itemsUrl = re.findall(urlpre, content.encode('utf8'))
    for item in itemsUrl:
        uPre = item[0]
        uNext = item[1]
    itemsUname = re.findall(nName, content.encode('utf8'))
    for item in itemsUname:
        uName = item
    # itemsNext = re.findall(urlNext, content.encode('utf8'))
    # for item in itemsNext:
    #    uNext = item

    # 去掉章节目录
    tit = re.sub(replaceBrIndex, "", tit)

    # 把空格键替换， 把<br / >换成回车
    arti = re.sub(replaceBr, "\n", arti)
    arti = re.sub(replaceSpace, " ", arti)

    retStr = tit + "-|-" + arti
    retStr = retStr + "-|-" + uPre + "-|-" + uNext + "-|-" + uName
    print retStr
    # print itemsUrl


except urllib2.URLError, e:
    if hasattr(e, "code"):
        print e.code
    if hasattr(e, "reason"):
        print e.reason
