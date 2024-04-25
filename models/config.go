package models

var GlobalConfig Config

type Config struct {
	Host    string   `mapstructure:"HOST"`
	Port    int      `mapstructure:"PORT"`
	ConnStr string   `mapstructure:"CONNECTION_STRING"`
	Cors    []string `mapstructure:"CORS"`
}
