package database

import (
	"address/config"
	"address/models"
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	// MysqlDB MysqlDB
	MysqlDB *gorm.DB
	// once once
	once sync.Once
	// err err
	err error
)

// InitAll InitAll
func init() {
	once.Do(func() {
		InitMysql()
	})

	MysqlDB.AutoMigrate(&models.Address{})
}

// InitMysql InitMysql
func InitMysql() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=%s",
		config.GetDefaultEnv("DB_USERNAME", "root"),
		config.GetDefaultEnv("DB_PASSWORD", "123456"),
		config.GetDefaultEnv("DB_HOST", "127.0.0.1"),
		config.GetDefaultEnv("DB_PORT", "3306"),
		config.GetDefaultEnv("DB_DATABASE", "database"),
		strings.Replace(config.GetDefaultEnv("DB_TIMEZONE", "Local"), "/", "%2F", -1),
	)
	MysqlDB, err = gorm.Open("mysql", dsn)
	if err != nil {
		log.Panicf("init DB err:%v,dsn:%v", err, dsn)
	}
	//defer MysqlDB.Close()
	// 全局禁用表名复数
	MysqlDB.SingularTable(true)
	//启用Logger，显示详细日志，默认情况下会打印发生的错误
	//Conn.LogMode(false)
	MysqlDB.DB().SetMaxIdleConns(100)
	MysqlDB.DB().SetMaxOpenConns(100)

}

// CloseMysql 关闭mysql
func CloseMysql() {
	defer MysqlDB.Close()
}
