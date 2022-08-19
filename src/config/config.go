package config

type Config struct {
	Server  *Server  `yaml:"Server"`
	Logger  *Logger  `yaml:"Logger"`
	WebHook *WebHook `yaml:"WebHook"`
}

type Logger struct {
	LogDir   string `yaml:"LogDir"`   // 日志存储路径
	LogLevel string `yaml:"LogLevel"` // 日志默认级别，值必须是：debug,info,warn,error，默认值是info
}

type Server struct {
	Address  string `yaml:"Address"`  // 监听地址
	Port     uint64 `yaml:"Port"`     // 监听端口
	Context  string `yaml:"Context"`  // 上下文路径
	LogDir   string `yaml:"LogDir"`   // 日志存储路径
	LogLevel string `yaml:"LogLevel"` // 日志默认级别，值必须是：debug,info,warn,error，默认值是info
}

type WebHook struct {
	ContentType string `yaml:"ContentType"` // 默认为 application/json , 若是旧版钩子(已不维护)为 application/x-www-form-urlencoded
	UserAgent   string `yaml:"UserAgent"`   // 固定为 git-oschina-hook，可用于标识为来自 gitee 的请求
	Token       string `yaml:"Token"`       // 用户新建 WebHook 时提供的密码或根据提供的签名密钥计算后的签名
}

type CMD struct {
	Port   uint64 // 服务端口
	Sh     string // 指令
	File   string // 文件路径
	Ticker int64  // 定时器执行间隔
	Cron   string // 定时器执行文件
}
