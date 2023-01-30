package setting

import (
	"fmt"

	"github.com/spf13/viper"
)

type AppConfig struct {
	Mode         string `mapstructure:"mode"`
	Port         int    `mapstructure:"port"`
	Name         string `mapstructure:"name"`
	StartTime    string `mapstructure:"start_time"`
	MachineID    int    `mapstructure:"machine_id"`
	*LogConfig   `mapstructure:"log"`
	*MysqlConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
	*JwtConfig   `mapstructure:"JwtAuth"`
	*FtpConfig   `mapstructure:"FTP"`
	*VideoConfig `mapstructure:"Video"`
}

type LogConfig struct{}

type MysqlConfig struct {
	Host         string `mapstructure:"Host"`
	User         string `mapstructure:"User"`
	Pwd          string `mapstructure:"Pwd"`
	DB           string `mapstructure:"DBname"`
	Port         string `mapstructure:"Port"`
	MaxOpenConns int    `mapstructure:"Max_open_conns"`
	MaxIdleConns int    `mapstructure:"Max_idle_conns"`
}

type RedisConfig struct{}

type JwtConfig struct {
	AccessExpire  string `mapstructure:"AccessExpire"`
	RefreshExpire string `mapstructure:"RefreshExpire"`
	Issuer        string `mapstructure:"Issuer"`
}

type FtpConfig struct {
	ServerAddr string `mapstructure:"ServerAddr"`
	UserName   string `mapstructure:"Name"`
	Pwd        string `mapstructure:"Pwd"`
}

type VideoConfig struct {
	PlayUrlPrefix  string `mapstructure:"PlayUrlPrefix"`
	CoverUrlPrefix string `mapstructure:"CoverUrlPrefix"`
}

var Conf = new(AppConfig)

func Init() error {
	viper.SetConfigFile("./config/config.yaml")
	viper.WatchConfig()
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("ReadInConfig failed, err: %v", err))
	}
	if err := viper.Unmarshal(&Conf); err != nil {
		panic(fmt.Errorf("unmarshal to Conf failed, err:%v", err))
	}
	return err
}
