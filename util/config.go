package util

import (
	"fmt"
	"github.com/spf13/viper"
	"time"
)

// Config stores all configuration of the application.
// The values are read by viper from a config file or environment variable.
type Config struct {
	AppReadTimeout  time.Duration `mapstructure:"APP_READ_TIMEOUT"`
	AppWriteTimeout time.Duration `mapstructure:"APP_WRITE_TIMEOUT"`

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
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", c.DBUser, c.DBPassword, c.DBHost, c.DBPort, c.DBName, c.SSLMode)
}

// LoadConfig reads configuration from environment file or variables
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
