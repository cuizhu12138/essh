package data

import (
	"essh/config"
	"fmt"

	"github.com/spf13/viper"
)

// 定义各项数据
type Target struct {
	User        string
	Pwd         string
	Host        Host
	Description string
}

// 定义Host结构
type Host struct {
	Address string
	Port    int
}

var (
	// 所有主机的列表
	HostList []Target
)

// 读取targetlist.json中的数据 进行初始化
func InitHost() {
	// Set Default Value
	viper.SetDefault("user", "root")
	viper.SetDefault("pwd", "123456")
	viper.SetDefault("host", Host{
		Address: "127.0.0.1",
		Port:    22,
	})
	viper.SetDefault("description", "nil")

	// Set config file message
	viper.SetConfigFile("./targetlist.json")
	viper.SetConfigName("targetlist")
	viper.SetConfigType("json")
	viper.AddConfigPath("./")

	// 尚未考虑文件不存在的情况
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("error in host.go: %w", err))
	}

	// debug
	if config.DebugMode {
		fmt.Println("init host list successfully")
	}

	// 排除错误的json文件
	allhost := viper.Get("allhost")
	if allhost == nil {
		panic("error in host.go: Error In Json")
	}

	// 解析错误
	err := viper.UnmarshalKey("allhost", &HostList)
	if err != nil {
		panic(fmt.Errorf("error in host.go: %w", err))
	}
}
