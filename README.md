# GoReadNovel

#2017年01月16日
构建框架


#2017年01月18日
加入爬虫


# 下载小说的请求


id 是最后一位  和文件名字
/modules/article/txtarticle.php?id=11&fname=圣墟

# 添加json


GetTopNovelListJson  
GetBookContentJson 
GetSearchNovelJson
GetNovelInfoJson   ?go=/book/0/11/
GetTopByTypeNovelListJson  ntype stype page

#修改为https模式

添加server.crt,server.key

生成命令

openssl genrsa -out server.key 2048
openssl req -new -x509 -key server.key -out server.crt -days 365
