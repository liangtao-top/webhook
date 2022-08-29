# WebHook Git自动化部署辅助工具

- Http服务监听WebHook请求。
- 支持密码、签名验证。
- 支持配置化命令。
- 支持自定义定时器备份数据库，文件等。

## 安装／运行

### step 1: 拉取代码

~~~
git clone https://github.com/liangtao-top/webhook.git
~~~

### step 2: 编译启动

~~~
cd webhook
./dist/webhook.exe -sh "git status"
~~~

## 参数

~~~
  -p uint
        Http服务端口 (default 9527)
  -port uint
        Http服务端口 (default 9527)
  -context string
        Http服务上下文路径 (default /webhook)
  -cmd string
        WebHook预执行指令
  -sh string
        WebHook预执行指令
  -f string
        WebHook预执行文件
  -file string
        WebHook预执行文件
  -t uint
        定时器预执行间隔，单位：秒 (default 86400)
  -ticker uint
        定时器预执行间隔，单位：秒 (default 86400)
  -c string
        定时器指令，【脚本文件:执行间隔（单位秒）】 数据库备份示例：-c "/app/script/backup-mysql.sh:3600"
  -cron string
        定时器指令，多文件示例：-cron "/app/script/backup-mysql.sh:3600,/app/script/backup-www.sh:86400"
  -token string
        用户的 WebHook 密钥
~~~

## 常见问题

### 查询某个端口是否被占用

~~~
netstat -anp | grep 9527
~~~

### 终止进程

kill [参数] [进程号]

~~~
kill -9 进程号
~~~