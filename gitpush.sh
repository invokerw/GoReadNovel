#!/bin/bash


mysqldump -uroot -pweifei novel > /root/gopro/src/GoReadNovel/novel.sql

git status


git add -A


git commit -m "$1"


git push
