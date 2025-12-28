package config

import (
	"fmt"
	"strings"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type Config struct {
	Api ApiConfig
	Web WebConfig
	DB  DbConfig
}

func LoadConfig() (*Config, error) {
	var cfg Config

	setDefault()

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")
	viper.AddConfigPath("/etc/gomono/")

	_ = viper.ReadInConfig()

	viper.SetConfigType("env")
	viper.SetConfigFile(".env")

	if err := viper.MergeInConfig(); err != nil {
		fmt.Println("no .env file found")
	}

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	pflag.Int("api.port", 8080, "Api port")
	pflag.Parse()

	viper.BindPFlags(pflag.CommandLine)

	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func setDefault() {
	viper.SetDefault("api.port", 4000)
	viper.SetDefault("web.port", 4001)

	// db default
	viper.SetDefault("db.host", "localhost")
	viper.SetDefault("db.port", 5400)
	viper.SetDefault("db.user", "admin")
	viper.SetDefault("db.password", "p@sw0rd!")
	viper.SetDefault("db.name", "test_db")
	viper.SetDefault("db.url", "postgresql://admin:p@sw0rd!@127.0.0.1:5400/test_db?sslmode=disable")
}
