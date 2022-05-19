package config

import (
	"fmt"
)

var (
	def = Config{LogsAddress: "./logs/"}
)

const (
	configFormat = "AppConfig:%v {" +
		"\n\tPort: %d" +
		"\n\tMysql: %s" +
		"\n\tRedis: %s" +
		"\n\tLevel: %v" +
		"\n\tLogOut >> %v" +
		"\n}"

	mysqlFormat = "%s@%s:%d"
	redisFormat = "%s:%d"
)

func GetDefConf() Config {

	return def
}

type Config struct {
	Name        string      `mapstructure:"name"`
	Port        int         `mapstructure:"port"`
	Mysql       MysqlConfig `mapstructure:"mysql"`
	Redis       RedisConfig `mapstructure:"redis"`
	LogLevel    int         `mapstructure:"logLevel"`
	LogsAddress string      `mapstructure:"logs"`
}

// ToDeBug 配置信息
func (c Config) String() string {
	return fmt.Sprintf(configFormat, c.Name, c.Port, c.Mysql, c.Redis, c.LogLevel, c.LogsAddress)
}

type MysqlConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Name     string `mapstructure:"name"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbName"`
}

func (m MysqlConfig) String() string {
	return fmt.Sprintf(mysqlFormat, m.Name, m.Host, m.Port)
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Password string `mapstructure:"passwd"`
	Port     int    `mapstructure:"port"`
	No       int    `mapstructure:"no"`
}

func (r RedisConfig) String() string {
	return fmt.Sprintf(redisFormat, r.Host, r.Port)
}
