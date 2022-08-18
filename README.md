# WebHook Git自动化部署辅助工具

- Http服务监听WebHook请求。
- 支持密码、签名验证。
- 支持配置化命令

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

## 运行参数

~~~
  -port uint64
  -p uint64
        监听端口
  -sh
  -cmd string
        预执行指令
  -file string
  -f string
        预执行脚本文件路径
~~~