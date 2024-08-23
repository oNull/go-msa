package main

import (
	"encoding/json"
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
)

type SeverConfig struct {
	Name        string        `mapstructure:"name"`
	Port        int           `mapstructure:"port"`
	UserSrvInfo UserSrvConfig `mapstructure:"user_srv"`
	JWTInfo     JWTConfig     `mapstructure:"jwt"`
	JuHe        JHConfig      `mapstructure:"juhe"`
	RedisInfo   RedisConfig   `mapstructure:"redis"`
	ConsulInfo  ConsulConfig  `mapstructure:"consul"`
}

func main() {

	// 至少一个ServerConfig
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr:      "124.222.97.37",
			ContextPath: "/nacos",
			Port:        8848,
			Scheme:      "http",
		},
	}
	// 创建clientConfig
	clientConfig := constant.ClientConfig{
		NamespaceId:         "a97d1437-7b35-40c5-abc3-2a5e8b49f78f", // 如果需要支持多namespace，我们可以创建多个client,它们有不同的NamespaceId。当namespace是public时，此处填空字符串。
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "./tmp/nacos/log",
		CacheDir:            "./tmp/nacos/cache",
		LogLevel:            "debug",
	}

	// 创建clientConfig的另一种方式
	// 创建动态配置客户端的另一种方式 (推荐)
	configClient, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)

	if err != nil {
		fmt.Println("我是第一个错误")
		panic(err)
	}

	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: "user_web.json",
		Group:  "dev",
	})

	if err != nil {
		fmt.Println("我是第二个错误")
		panic(err)
	}

	fmt.Println(content)
	var ser SeverConfig
	err = json.Unmarshal([]byte(content), &ser)
	if err != nil {
		fmt.Printf("读取nacos配置失败： %s", err.Error())
		return
	}
	// 输出解析后的结构体数据
	fmt.Printf("%+v\n", ser)

	//// 将解析后的数据转换为JSON格式
	//jsonData, err := json.MarshalIndent(ser, "", "  ")
	//if err != nil {
	//	log.Fatalf("无法将数据转换为JSON: %v", err)
	//}
	//
	//// 输出转换后的JSON数据
	//fmt.Println(string(jsonData))
	// 还需要监听配置文件变化
}

type JWTConfig struct {
	SigningKey string `mapstructure:"key"`
}

type JHConfig struct {
	Url string `mapstructure:"url"`
	Key string `mapstructure:"key"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
}

type ConsulConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
	//Name string `mapstructure:"name"`
}

type UserSrvConfig struct {
	//Host string `mapstructure:"host"`
	//Port int    `mapstructure:"port"`
	Name string `mapstructure:"name"`
}
