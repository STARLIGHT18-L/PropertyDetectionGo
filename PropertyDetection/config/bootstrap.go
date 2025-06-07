package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/minio/minio-go"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"os"
)

// 初始化配置文件
func ConfigInit() *viper.Viper {
	config := "config.yaml"
	if configEnv := os.Getenv("VIPER_CONFIG"); configEnv != "" {
		config = configEnv
	}
	// 初始化配置文件
	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	//监听配置文件变化并热加载
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config file changed:", in.Name)
		if err := v.Unmarshal(&Boot.Config); err != nil { //重载配置
			fmt.Println("Boot.Config:", err)
		}
	})
	if err := v.Unmarshal(&Boot.Config); err != nil {
		fmt.Println("Boot.Config:", err)
	}
	return v
}

// 数据库初始化
func MysqlInit() *gorm.DB {
	switch Boot.Config.Database.Driver {
	case "mysql":
		return InitMysqlGorm()
	default:
		return InitMysqlGorm()
	}
}
func MinioInit() *minio.Client {
	minioClient, err := minio.New(Boot.Config.Minio.Host, Boot.Config.Minio.AccessKey, Boot.Config.Minio.SecretKey, false)
	if err != nil {
		panic(err)
	}
	MinioBucketInit(minioClient)
	return minioClient
}
