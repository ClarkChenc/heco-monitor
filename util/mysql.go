package util

import (
	"fmt"
	"net/url"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

var DB *gorm.DB

func InitDBconfig() {
	//制定配置文件的路径
	viper.SetConfigFile("conf/db.yaml")
	// 读取配置信息
	err := viper.GetViper().MergeInConfig()
	if err != nil {
		// 读取配置信息失败
		fmt.Printf("viper.ReadInConfig() failed,err:%v\n", err)
		return
	}
	//监听修改
	viper.WatchConfig()
	//为配置修改增加一个回调函数
	// viper.OnConfigChange(func(in fsnotify.Event) {
	// 	fmt.Println("配置文件修改了...")
	// })
	return
}

func InitDB() *gorm.DB {
	driverName := viper.GetString("monitorDB.driverName")
	host := viper.GetString("monitorDB.host")
	port := viper.GetString("monitorDB.port")
	database := viper.GetString("monitorDB.database")
	username := viper.GetString("monitorDB.username")
	password := viper.GetString("monitorDB.password")
	charset := viper.GetString("monitorDB.charset")
	loc := viper.GetString("monitorDB.loc")
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=%s",
		username,
		password,
		host,
		port,
		database,
		charset,
		url.QueryEscape(loc))
	var db *gorm.DB
	var err error

	for {
		db, err = gorm.Open(driverName, args)
		if err != nil {
			fmt.Println("fail to connect database, err: " + err.Error())
			time.Sleep(time.Second * 3)
			continue
		}
		break
	}

	fmt.Println("[SUCCESS] DB connected")

	DB = db
	return db
}

func GetDB() *gorm.DB {
	return DB
}

func CloseDB() {
	fmt.Println("[SUCCESS] DB Closed")
	DB.Close()
}
