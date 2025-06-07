package config

import (
	"github.com/minio/minio-go"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type Global struct {
	Viper  *viper.Viper
	Mysql  *gorm.DB
	Cache  *Cache
	Minio  *minio.Client
	Config Configuration
}

var Boot = new(Global)
