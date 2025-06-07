package config

import base "PropertyDetection/model/base"

type Configuration struct {
	App      base.App      `json:"app" yaml:"app" mapstructure:"app"`
	Jwt      base.Jwt      `json:"jwt" yaml:"jwt" mapstructure:"jwt"`
	Minio    base.Minio    `json:"minio" yaml:"minio" mapstructure:"minio"`
	Database base.Database `json:"database" yaml:"database" mapstructure:"database"`
}
