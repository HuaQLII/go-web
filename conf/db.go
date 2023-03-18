package conf

import (
	"github.com/spf13/viper"
	"github/im-lauson/go-web/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"time"
)

func InitDb() (*gorm.DB, error) {
	logMode := logger.Info
	if !viper.GetBool("mode.develop") {
		logMode = logger.Error
	}
	db, err := gorm.Open(mysql.Open(viper.GetString("db.dsn")), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "sys_",
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(logMode),
	})
	if err != nil {
		return nil, err
	}
	sqlDb, _ := db.DB()
	sqlDb.SetMaxIdleConns(viper.GetInt("db.maxIdleConn"))    //最大空闲连接数
	sqlDb.SetMaxOpenConns(viper.GetInt("db.setMaxOpenConn")) //最大连接数
	sqlDb.SetConnMaxLifetime(time.Hour)                      //最大连接时常
	db.AutoMigrate(&model.User{})

	return db, nil
}
