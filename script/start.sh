#!/bin/bash
./dist/webhook-linux -sh "git --version" -file "/app/script/git/git-tag-develop.sh" -ticker 3600 -cron "/app/script/dev/backup-mysql.sh"
