#!/bin/bash
nohup ./dist/webhook-linux -file "xxx.sh" -token "xxx" -cron "backup-xxx.sh:3600" >/dev/null 2>&1 &