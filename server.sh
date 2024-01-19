#!/bin/bash
cmd=$2
usage="Usage: $(basename $0) (up|stop|restart|status)"

#webhook脚本路径参数
webhook_path=/app/script/bin/webhook.sh

#没有输入参数时提醒内容 $#参数的个数
if [ $# -eq 0 ]; then
  echo "please input up|stop|restart|status"
  exit
fi

# 启动时执行的命令
start() {
  $webhook_path $cmd
}

function logs() {
  for dir in $(ls $1); do
    if [ -d $1"/"$dir ]; then
      for file in $(ls -t $1"/"$dir); do
        log=$1"/"$dir"/"$file
        echo $log
        tail -f $log &
        break 2
      done
    fi
  done
}

#启动服务
up() {
  vehicle_repid=$(ps -ef | grep webhook-linux | grep -v 'grep' | awk '{print $2}')
  if [ "$vehicle_repid" ]; then
    #    echo "$(ps -ef | grep webhook-linux | grep -v 'grep')"
    #    echo "服务已启动"
    logs /webhook/logs
  else
    start
    if [ "$cmd" == "-d" ]; then
      vehicle_pid_temp=$(ps -ef | grep webhook-linux | grep -v 'grep' | awk '{print $2}')
      if [ "$vehicle_pid_temp" ]; then
        echo "$(ps -ef | grep webhook-linux | grep -v 'grep')"
        echo "webhook startup suessful!"
      else
        echo "服务启动失败"
      fi
    fi
  fi
}

#杀死服务 注意 [] 和$pid 存在空格
stop() {
  #获取需要杀死进程的id
  pid=$(ps -ef | grep webhook-linux | grep -v 'grep' | awk '{print $2}')
  if [ "$pid" ]; then
    for id in $pid; do
      echo "kill -9 $id"
      kill -9 $id
    done
    echo "webhook stop suessful!"
  else
    echo "webhook pid is not exists"
  fi
}

#重启服务
restart() {
  stop
  up
}

case $1 in
up) #启动进程
  up
  ;;
stop) #根据输入命令判断需要执行的操作
  stop
  ;;
restart) #重启程序
  restart
  ;;
status) #查询状态
    vehicle_repid=$(ps -ef | grep webhook-linux | grep -v 'grep' | awk '{print $2}')
    if [ "$vehicle_repid" ]; then
        echo "$(ps aux | grep webhook-linux | grep -v 'grep')"
    else
        echo "webhook service not running!"
    fi
  ;;
*)
  echo "Error command"
  echo "$usage"
  ;;
esac
