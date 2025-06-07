package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

// 获取Gorm日志写入器
func GetGormLogWriter() logger.Writer {
	var writer io.Writer
	// 判断是否启用文件日志写入
	if Boot.Config.Database.EnableFileLogWriter {
		//  文件日志(这里用打印替代)
		writer = os.Stdout
	} else {
		writer = os.Stdout
	}
	return log.New(writer, "\r\n", log.LstdFlags)
}

// 获取Gorm日志
func GetGormLogger() logger.Interface {
	var logMode logger.LogLevel
	switch Boot.Config.Database.LogMode {
	case "silent":
		logMode = logger.Silent
	case "error":
		logMode = logger.Error
	case "warn":
		logMode = logger.Warn
	case "info":
		logMode = logger.Info
	default:
		logMode = logger.Info
	}
	return logger.New(GetGormLogWriter(), logger.Config{
		SlowThreshold:             500 * time.Millisecond,                    //  慢SQL阈值
		LogLevel:                  logMode,                                   // 日志级别
		IgnoreRecordNotFoundError: false,                                     //  记录未找到的错误
		Colorful:                  !Boot.Config.Database.EnableFileLogWriter, //  是否禁用彩色打印
	})
}

// 数据库表初始化
func InitMySqlTables(db *gorm.DB) {
	//err := db.AutoMigrate(
	//)
	//if err != nil {
	//	Boot.Log.Error("migrate table failed", zap.Any("err", err))
	//	os.Exit(0)
	//}
}
func InitMysqlGorm() *gorm.DB {
	dbConfig := Boot.Config.Database
	if dbConfig.Database == "" {
		return nil
	}
	dsn := dbConfig.UserName +
		":" + dbConfig.Password +
		"@tcp(" + dbConfig.Host +
		":" + strconv.Itoa(dbConfig.Port) + ")/" +
		dbConfig.Database +
		"?charset=" +
		dbConfig.Charset +
		"&parseTime=True&loc=Local"
	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         255,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,            // 禁用自动创建外键约束
		Logger:                                   GetGormLogger(), // 使用自定义 Logger
	}); err != nil {
		fmt.Println("mysql connect error:", err)
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(dbConfig.MaxIdleConns)
		sqlDB.SetMaxOpenConns(dbConfig.MaxOpenConns)
		InitMySqlTables(db)
		return db
	}
}
