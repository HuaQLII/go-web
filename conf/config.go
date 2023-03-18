package conf

import (
	"fmt"
	"github.com/spf13/viper"
)

func InitConfig() {
	//读取相关的配置文件
	viper.SetConfigName("settings")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./conf/")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Sprintf("Load Config Error:%s", err.Error()))
	}

	fmt.Println("==========================")
	fmt.Println(viper.GetString("server.port"))
	fmt.Println("==========================")
}
