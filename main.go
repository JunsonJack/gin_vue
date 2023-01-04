package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"junsonjack.cn/go_vue/common"
)

func main() {
	// 读取配置文件
	InitConfig()

	// 初始化数据库连接
	db := common.InitDB()
	defer db.Close() //延时关闭连接

	r := gin.Default()
	r = CollectRoute(r)

	// 修改端口
	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port))
	}

	panic(r.Run())
	
	// panic(r.Run(":8088"))
}
// 使用viper管理配置文件
func InitConfig (){
	workDir,_ := os.Getwd() //获取工作目录
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

