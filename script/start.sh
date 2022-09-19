#!/bin/bash
nohup ./dist/webhook-linux -file "xxx.sh" -cron "/xxx.sh" -ticker 3600 >/dev/null 2>&1 &