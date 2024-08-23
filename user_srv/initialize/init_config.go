package initialize

import (
	"encoding/json"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"user_srv/global"
)

func GetEnvInfo(env string) string {
	viper.AutomaticEnv()
	return viper.GetString(env)
	//刚才设置的环境变量 想要生效 我们必须得重启goland
}

func InitConfig() {
	zap.S().Info("初始化配置文件...")

	data := GetEnvInfo("DEBUG")
	var configFileName string
	configFileNamePrefix := "config"
	if data == "true" {
		configFileName = fmt.Sprintf("./config/%s-debug.yaml", configFileNamePrefix)
	} else {
		configFileName = fmt.Sprintf("./config/%s-pro.yaml", configFileNamePrefix)
	}

	v := viper.New()
	v.SetConfigFile(configFileName)
	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}

	//这个对象如何在其他文件中使用 - 全局变量
	if err := v.Unmarshal(global.NacosConfig); err != nil {
		panic(err)
	}
	zap.S().Infof("配置信息: %v", global.NacosConfig)

	//从nacos中读取配置信息
	sc := []constant.ServerConfig{
		{
			IpAddr: global.NacosConfig.Host,
			Port:   global.NacosConfig.Port,
		},
	}

	cc := constant.ClientConfig{
		NamespaceId:         global.NacosConfig.Namespace, // 如果需要支持多namespace，我们可以场景多个client,它们有不同的NamespaceId
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "tmp/nacos/log",
		CacheDir:            "tmp/nacos/cache",
		LogLevel:            "debug",
	}

	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": sc,
		"clientConfig":  cc,
	})
	if err != nil {
		panic(err)
	}

	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: global.NacosConfig.DataId,
		Group:  global.NacosConfig.Group})

	if err != nil {
		panic(err)
	}
	//fmt.Println(content) //字符串 - yaml
	//想要将一个json字符串转换成struct，需要去设置这个struct的tag
	err = json.Unmarshal([]byte(content), &global.ServerConfig)
	if err != nil {
		zap.S().Fatalf("读取nacos配置失败： %s", err.Error())
	}
	fmt.Println(&global.ServerConfig)
}

func InitConfig2() {
	data := GetEnvInfo("DEBUG")
	zap.S().Infof("获取ENV：%s", data)
	var configFileName string
	configFileNamePrefix := "config"
	if data == "true" {
		zap.S().Infof("获取测试配置")
		configFileName = fmt.Sprintf("./config/%s-debug.yaml", configFileNamePrefix)
	} else {
		zap.S().Infof("获取生产配置")
		configFileName = fmt.Sprintf("./config/%s-pro.yaml", configFileNamePrefix)
	}

	v := viper.New()
	v.SetConfigFile(configFileName)
	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}
	if err := v.Unmarshal(&global.ServerConfig); err != nil {
		panic(err)
	}

	zap.S().Infof("配置文件：%s,配置信息：%v", configFileName, global.ServerConfig)
	v.OnConfigChange(func(e fsnotify.Event) {
		zap.S().Infof("配置文件产生变化：%v", e.Name)
		_ = v.ReadInConfig() // 读取配置数据
		_ = v.Unmarshal(global.ServerConfig)
		zap.S().Infof("配置信息为：%v", global.ServerConfig)
	})
}
