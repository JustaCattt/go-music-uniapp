package db

import (
	"Music/server/model"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/viper"
	"time"
)

type DbConfig struct {
	Driver, URL string
}

var PGEngine *gorm.DB

//初始化postgres数据库
func InitPgDB() *gorm.DB {
	dbConfig := ConfigParser()                    //解析赋值
	db := NewConnection(dbConfig)                 //连接数据库
	db.DB().SetMaxIdleConns(10)                   //最大空闲连接数
	db.DB().SetMaxOpenConns(30)                   //最大连接数
	db.DB().SetConnMaxLifetime(time.Second * 300) //设置连接空闲超时
	db.SingularTable(true)                        //如果使用gorm来帮忙创建表时，这里填写false的话gorm会给表添加s后缀，填写true则不会
	db.LogMode(true)                              //打印sql语句
	db.AutoMigrate(&model.User{})                 //自动迁移建表，实时更新表
	PGEngine = db                                 //赋值给PGEngine，供外部使用
	return db
}

//连接数据库
func NewConnection(dbConfig DbConfig) *gorm.DB {
	conn, err := gorm.Open(dbConfig.Driver, dbConfig.URL)
	if err != nil {
		panic("failed to connect database:\t"+err.Error())
	}
	return conn
}

//获取数据库配置信息
func ConfigParser() (dbConfig DbConfig) {
	dbConfig.Driver = viper.GetString("datasource.driver")
	url := viper.GetString("datasource.url")
	if url != "" {
		dbConfig.URL = url
	}else {
		user := viper.GetString("datasource.user")
		password := viper.GetString("datasource.password")
		host := viper.GetString("datasource.host")
		dbname := viper.GetString("datasource.database")
		port := viper.GetInt("datasource.port")
		if host == "" {
			host = "127.0.0.1"
		}
		//拼接连接数据库的url
		switch dbConfig.Driver {
		case "postgres":
			if port == 0 {
				port = 5432
			}
			dbConfig.URL = fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=disable", user, password, host, port, dbname)
		case "mysql":
			if port == 0 {
				port = 3306
			}
			dbConfig.URL = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, port, dbname)
		}
	}
	return dbConfig
}
