package configs

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

func init() {
	viper.SetConfigName("config") // 读取yaml配置文件
	viper.AddConfigPath("./configs")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			log.Println("no such config file")
		} else {
			// Config file was found but another error was produced
			log.Println("read config error")
		}
		log.Fatal(err) // 读取配置文件失败致命错误
	}
}

func GetStringConfig(key string) string {
	//fmt.Println("获取了", key)
	return viper.GetString(key)
}
func GetIntConfig(key string) int {
	return viper.GetInt(key)
}

func Test() {
	fmt.Println("获取配置文件的mysql.username", GetStringConfig("mysql.username"))
}
