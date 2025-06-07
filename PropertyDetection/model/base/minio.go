package base

type Minio struct {
	Host      string `mapstructure:"host" json:"host" yaml:"host"`
	Bucket    string `mapstructure:"bucket" json:"bucket" yaml:"bucket"`
	AccessKey string `mapstructure:"access_key" json:"access_key" yaml:"access_key"`
	SecretKey string `mapstructure:"secret_key" json:"secret_key" yaml:"secret_key"`
	Url       string `mapstructure:"url" json:"url" yaml:"url"`
}
