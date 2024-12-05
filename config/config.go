package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
)

type (
	Config struct {
		Env          string `mapstructure:"env"`
		Server       `mapstructure:"http-server"`
		MusicInfoUrl string `mapstructure:"music-info-url"`
		Db           `mapstructure:"database"`
	}
	Server struct {
		Address string
	}
	Db struct {
		Dbname   string
		Username string
		Host     string
		Port     string
		Password string
		Sslmode  string
		Dsn      string `mapstructure:"dsn"`
	}
)

func GetConfig() *Config {
	var config *Config
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	viper.AutomaticEnv()
	configPath := viper.GetString("CONFIG_PATH")
	if configPath == "" {
		log.Fatalf("CONFIG_PATH environment variable is not set")
	} else {
		viper.AddConfigPath(configPath)
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
	}
	err = viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error while reading configs file %v", err)
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("Unmarshal error: %v", err)
	}
	config.Db = Db{
		Dbname:   viper.GetString("POSTGRES_DB"),
		Username: viper.GetString("POSTGRES_USER"),
		Host:     viper.GetString("POSTGRES_HOST"),
		Port:     viper.GetString("POSTGRES_PORT"),
		Password: viper.GetString("POSTGRES_PASSWORD"),
		Sslmode:  viper.GetString("POSTGRES_SSLMODE"),
	}
	config.Db.Dsn = fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s", config.Db.Host, config.Db.Port, config.Db.Dbname, config.Db.Username, config.Db.Password, config.Db.Sslmode)
	return config
}
