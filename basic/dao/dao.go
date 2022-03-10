package dao

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var writeDbSource *gorm.DB
var readDbSource *gorm.DB

// BaseDao 基础dao
type BaseDao struct {
}

// GetReadDbSource 获取读库source
func (dao *BaseDao) GetReadDbSource() *gorm.DB {
	return readDbSource
}

// GetWriteDbSource 获取写库source
func (dao *BaseDao) GetWriteDbSource() *gorm.DB {
	return writeDbSource
}

// InitializeDao 初始化dao
func InitializeDao() {
	// 为了处理time.Time，您需要包括parseTime作为参数。
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=%t&loc=%s",
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.host"),
		viper.GetInt("db.port"),
		viper.GetString("db.dbName"),
		true,
		"Local")

	var err error
	writeDbSource, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "t_",
			SingularTable: true,
		},
	})
	if err != nil {
		panic(err)
	}

	sqlDB, err := writeDbSource.DB()
	sqlDB.SetMaxIdleConns(20)
	sqlDB.SetMaxOpenConns(60)
	sqlDB.SetConnMaxLifetime(time.Hour)

	readDbSource, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "t_",
			SingularTable: true,
		},
	})
	if err != nil {
		panic(err)
	}

	sqlDB, err = readDbSource.DB()
	sqlDB.SetMaxIdleConns(50)
	sqlDB.SetMaxOpenConns(200)
	sqlDB.SetConnMaxLifetime(time.Hour)
}
