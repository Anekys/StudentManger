package utils

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

func getStringConfig(key string) string {
	var v = viper.New()
	v.SetConfigName("config") // 读取yaml配置文件
	v.AddConfigPath(".")

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
	fmt.Println("获取了", key)
	return v.GetString(key)
}
func getIntConfig(key string) int {
	var v = viper.New()
	return v.GetInt(key)
}

func Test() {
	fmt.Println("获取配置文件的mysql.username", getStringConfig("mysql.username"))
}
