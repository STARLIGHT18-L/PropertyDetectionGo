package base

type App struct {
	Env     string `json:"env" yaml:"env" mapstructure:"env"`
	Port    int    `json:"port" yaml:"port" mapstructure:"port"`
	AppName string `json:"app_name" yaml:"app_name" mapstructure:"app_name"`
	AppUrl  string `json:"app_url" yaml:"app_url" mapstructure:"app_url"`
}
