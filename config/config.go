package config

type Config struct {
	Name        string      `mapstructure:"name"`
	Port        int         `mapstructure:"port"`
	Mysql       MysqlConfig `mapstructure:"mysql"`
	Redis       RedisConfig `mapstructure:"redis"`
	LogsAddress string      `mapstructure:"logsAddress"`
}

type MysqlConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Name     string `mapstructure:"name"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbName"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Password int    `mapstructure:"passwd"`
	Port     int    `mapstructure:"port"`
}
