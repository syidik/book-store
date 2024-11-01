package config

import (
	"errors"
	"fmt"

	"github.com/spf13/viper"
)

var AppConfig Config

type Config struct {
	Port        int    `mapstructure:"PORT"`
	Environment string `mapstructure:"ENVIRONMENT"`
	Debug       bool   `mapstructure:"DEBUG"`

	PostgresHost string `mapstructure:"POSTGRES_HOST"`
	PostgresPort int    `mapstructure:"POSTGRES_PORT"`
	PostgresDB   string `mapstructure:"POSTGRES_DB"`
	PostgresUser string `mapstructure:"POSTGRES_USER"`
	PostgresPwd  string `mapstructure:"POSTGRES_PWD"`
}

func InitAppConfig() error {
	viper.SetConfigName(".env") // allow directly reading from .env file
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.AddConfigPath("internal/config")
	viper.AddConfigPath("/")
	viper.SetEnvPrefix("app")
	// viper.AutomaticEnv()
	err := viper.ReadInConfig()

	if err != nil {
		fmt.Print("failed to reading env file\n proceed binding from env")

		viper.BindEnv("Port", "PORT")
		viper.BindEnv("Environment", "ENVIRONMENT")
		viper.BindEnv("Debug", "DEBUG")

		viper.BindEnv("PostgresHost", "POSTGRES_HOST")
		viper.BindEnv("PostgresPort", "POSTGRES_PORT")
		viper.BindEnv("PostgresDB", "PostgresDB")
		viper.BindEnv("PostgresUser", "POSTGRES_USER")
		viper.BindEnv("PostgresPwd", "POSTGRES_PWD")
	}

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("internal/config")
	viper.AddConfigPath("/")
	viper.SetEnvPrefix("app")
	// viper.AutomaticEnv()
	errConfig := viper.MergeInConfig()

	if errConfig != nil {
		fmt.Print("failed to reading config file\n")
	}

	err = viper.Unmarshal(&AppConfig)
	if err != nil {
		return errors.New("failed to parse env to config struct\n")
	}

	if AppConfig.Port == 0 || AppConfig.PostgresHost == "" {
		return errors.New("required variabel environment is empty")
	}

	switch AppConfig.Environment {
	case "development":
		if AppConfig.PostgresHost == "" {
			return errors.New("required variabel environment is empty")
		}
	case "production":
		if AppConfig.PostgresHost == "" {
			return errors.New("required variabel environment is empty")
		}
	}

	return nil
}
