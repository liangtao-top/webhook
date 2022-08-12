package util

import (
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"webhook/src/global/enum"
	"webhook/src/logger"
)

func IniNacos2() {
	// 创建serverConfig
	serverConfigs := []constant.ServerConfig{{
		ContextPath: enum.CONFIG.Nacos.Server.ContextPath, // Nacos的ContextPath
		IpAddr:      enum.CONFIG.Nacos.Server.IpAddr,      // Nacos的服务地址
		Port:        enum.CONFIG.Nacos.Server.Port,        // Nacos的服务端口
		Scheme:      enum.CONFIG.Nacos.Server.Scheme,      // Nacos的服务地址前缀
	}}
	// 创建clientConfig
	clientConfig := ClientConfig()
	// 创建服务发现客户端
	namingClient, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)
	if err != nil {
		logger.Error("创建Nacos2服务发现客户端异常\n", err)
		return
	}

	ip := LocalIP()
	port := enum.CONFIG.Server.Port

	// 服务注册
	success, err := namingClient.RegisterInstance(vo.RegisterInstanceParam{
		Ip:          ip,
		Port:        port,
		ServiceName: "go_service_crawler",
		Weight:      10,
		Enable:      true,
		Healthy:     true,
		Ephemeral:   true,
		Metadata:    map[string]string{},
		ClusterName: "DEFAULT",       // 默认值DEFAULT
		GroupName:   "DEFAULT_GROUP", // 默认值DEFAULT_GROUP
	})
	if err != nil {
		logger.Error("Nacos2服务注册异常\n", err)
		return
	}
	var result string
	if success {
		result = "success"
	} else {
		result = "fail"
	}
	logger.Infof("Nacos2.0 swoole_service_gateway %s:%+v registry %+v.", ip, port, result)
}

func ClientConfig() constant.ClientConfig {
	return constant.ClientConfig{
		NamespaceId:         enum.CONFIG.Nacos.Client.NamespaceId,
		TimeoutMs:           enum.CONFIG.Nacos.Client.TimeoutMs,
		NotLoadCacheAtStart: enum.CONFIG.Nacos.Client.NotLoadCacheAtStart,
		LogDir:              enum.CONFIG.Nacos.Client.LogDir,
		CacheDir:            enum.CONFIG.Nacos.Client.CacheDir,
		LogLevel:            enum.CONFIG.Nacos.Client.LogLevel,
	}
}
