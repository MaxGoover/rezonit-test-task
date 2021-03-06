package util

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	DBDriver   string `mapstructure:"DB_DRIVER"`
	DBHost     string `mapstructure:"DB_HOST"`
	DBName     string `mapstructure:"DB_NAME"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBUser     string `mapstructure:"DB_USER"`
	SSLMode    string `mapstructure:"SSL_MODE"`

	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
}

func (c *Config) DBSource() string {
	return fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=%v", c.DBUser, c.DBPassword, c.DBDriver, c.DBPort, c.DBName, c.SSLMode)
}

func (c *Config) DBSourceTest() string {
	return fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=%v", c.DBUser, c.DBPassword, c.DBHost, c.DBPort, c.DBName, c.SSLMode)
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
