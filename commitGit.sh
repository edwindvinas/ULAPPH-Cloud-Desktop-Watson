#!/bin/bash
# commitGit.sh 202004120946PM
rm *.exe
rm *.exe~
git add --all
#git commit -m "update"
git commit -m $1
git push origin master

