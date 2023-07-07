#!/bin/bash

if [ "$1" == "-d" ]; then
  nohup ./dist/webhook-linux -p 9527 -context "/khzkYETe6gWv5eC1" -file "/app/script/git/git-tag-develop.sh" -token "u25839Vmx1Fu71mJUIkD1eaAOIeScxDx" -cron "/app/script/backup/backup-mysql.sh:604800,/app/script/backup/backup-www.sh:86400" >/dev/null 2>&1 &
  nohup ./dist/webhook-linux -p 9528 -context "/khzkYETe6gWv5eC1" -file "/web/script/git-tag-develop.sh" -token "u25839Vmx1Fu71mJUIkD1eaAOIeScxDx" >/dev/null 2>&1 &
else
  ./dist/webhook-linux -p 9527 -context "/khzkYETe6gWv5eC1" -file "/app/script/git/git-tag-develop.sh" -token "u25839Vmx1Fu71mJUIkD1eaAOIeScxDx" -cron "/app/script/backup/backup-mysql.sh:604800,/app/script/backup/backup-www.sh:86400" &
  ./dist/webhook-linux -p 9528 -context "/khzkYETe6gWv5eC1" -file "/web/script/git-tag-develop.sh" -token "u25839Vmx1Fu71mJUIkD1eaAOIeScxDx"
fi
