#!/bin/bash



if [ ! -n "$1" ] ;then
	echo "you have not input a commit word!"
	exit 1
else
	echo "the commit you input is $1"
	
	echo "Output Database Novel"

	mysqldump -uroot -pweifei novel > /root/gopro/src/GoReadNovel/novel.sql

	echo "OK"
	
	git status
	
	git add -A

	git commit -m "$1"

	git push
fi

exit 0


