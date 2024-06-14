package models

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"huiyi/config"
)

var DB *gorm.DB

// InitDB 初始化数据库，返回数据对象
func InitDB() error {
	var err error
	dbConfig := config.Config.Db

	// 构建数据源名称 (DSN)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&charset=%s",
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Db,
		dbConfig.Charset,
	)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Error("数据库连接失败: ", err)
		return err
	}

	// 检查是否连接成功
	sqlDB, err := DB.DB()
	if err != nil {
		log.Error("获取数据库连接对象失败: ", err)
		return err
	}

	// 测试数据库连接
	err = sqlDB.Ping()
	if err != nil {
		log.Error("数据库连接失败: ", err)
		return err
	}

	// 连接成功打印
	log.Info("数据库连接成功")
	return nil
}

// GetDB 返回数据库连接
func GetDB() *gorm.DB {
	return DB
}
