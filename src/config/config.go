package config

type Config struct {
	Server  *Server  `yaml:"Server"`
	Logger  *Logger  `yaml:"Logger"`
	Nacos   *Nacos   `yaml:"Nacos"`
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

type Nacos struct {
	Server *NacosServer `yaml:"Server"`
	Client *NacosClient `yaml:"Client"`
}

type NacosServer struct {
	ContextPath string `yaml:"ContextPath"` // Nacos的ContextPath
	IpAddr      string `yaml:"IpAddr"`      // Nacos的服务地址
	Port        uint64 `yaml:"Port"`        // Nacos的服务端口
	Scheme      string `yaml:"Scheme"`      // Nacos的服务地址前缀
}

type NacosClient struct {
	TimeoutMs   uint64 `yaml:"TimeoutMs"`   // 请求Nacos服务端的超时时间，默认是10000ms
	NamespaceId string `yaml:"NamespaceId"` // ACM的命名空间Id
	AppName     string `yaml:"AppName"`     // App名称
	Endpoint    string `yaml:"Endpoint"`    // 当使用ACM时，需要该配置. https://help.aliyun.com/document_detail/130146.html
	RegionId    string `yaml:"RegionId"`    // ACM&KMS的regionId，用于配置中心的鉴权
	AccessKey   string `yaml:"AccessKey"`   // ACM&KMS的AccessKey，用于配置中心的鉴权
	SecretKey   string `yaml:"SecretKey"`   // ACM&KMS的SecretKey，用于配置中心的鉴权
	OpenKMS     bool   `yaml:"OpenKMS"`     // 是否开启kms，默认不开启，kms可以参考文档 https://help.aliyun.com/product/28933.html
	// 同时DataId必须以"cipher-"作为前缀才会启动加解密逻辑
	CacheDir             string `yaml:"CacheDir"`             // 缓存service信息的目录，默认是当前运行目录
	UpdateThreadNum      int    `yaml:"UpdateThreadNum"`      // 监听service变化的并发数，默认20
	NotLoadCacheAtStart  bool   `yaml:"NotLoadCacheAtStart"`  // 在启动的时候不读取缓存在CacheDir的service信息
	UpdateCacheWhenEmpty bool   `yaml:"UpdateCacheWhenEmpty"` // 当service返回的实例列表为空时，不更新缓存，用于推空保护
	Username             string `yaml:"Username"`             // Nacos服务端的API鉴权Username
	Password             string `yaml:"Password"`             // Nacos服务端的API鉴权Password
	LogDir               string `yaml:"LogDir"`               // 日志存储路径
	LogLevel             string `yaml:"LogLevel"`             // 日志默认级别，值必须是：debug,info,warn,error，默认值是info
}

type CMD struct {
	Port uint64 // 服务端口
	Sh   string // 指令
	File string // 文件路径
}
