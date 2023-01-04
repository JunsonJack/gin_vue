package common

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"junsonjack.cn/go_vue/model"
)

var DB *gorm.DB

// 初始化数据库连接
func InitDB() *gorm.DB{
	driverName := viper.GetString("datasource.driverName")
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	database := viper.GetString("datasource.database")
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	charset := viper.GetString("datasource.charset")
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True",
			username,
			password,
			host,
			port,
			database,
			charset)
	// dsn := "root:123456@tcp(127.0.0.1:3306)/go_db?charset=utf8mb4&parseTime=True"
	db, err := gorm.Open(driverName, args)
	if err!=nil {
		panic("failed to connect database , err:" + err.Error())
	}

	// Migrate the schema
  db.AutoMigrate(&model.User{})

	DB = db
	return db
}

// 获取DB实例
func GetDB () *gorm.DB{
	return DB
}