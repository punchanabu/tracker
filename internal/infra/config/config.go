package config

import "github.com/spf13/viper"

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	Cache    CacheConfig
}

type ServerConfig struct {
	Port         string `mapstructure:"SERVER_PORT"`
	ReadTimeout  int    `mapstructure:"SERVER_READ_TIMEOUT"`
	WriteTimeout int    `mapstructure:"SERVER_WRITE_TIMEOUT"`
}

type DatabaseConfig struct {
	Host     string `mapstructure:"DB_HOST"`
	Port     string `mapstructure:"DB_PORT"`
	User     string `mapstructure:"DB_USER"`
	Password string `mapstructure:"DB_PASSWORD"`
	Name     string `mapstructure:"DB_NAME"`
	SSLMode  string `mapstructure:"DB_SSL_MODE"`
}

type CacheConfig struct {
	RedisURL      string `mapstructure:"REDIS_URL"`
	RedisPassword string `mapstructure:"REDIS_PASSWORD"`
}

func LoadConfig() (*Config, error) {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	// Load config into struct
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
