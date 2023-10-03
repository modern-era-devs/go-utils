package main

import (
	"fmt"
	"github.com/modern-era-devs/go-utils/config"
	"github.com/modern-era-devs/go-utils/database/postgres"
	"github.com/spf13/viper"
)

type AppConfig struct {
	Port                 int    `mapstructure:"PORT"`
	DBName               string `mapstructure:"DB_NAME"`
	DBPort               int    `mapstructure:"DB_PORT"`
	DBUser               string `mapstructure:"DB_USER"`
	DBPool               int    `mapstructure:"DB_POOL"`
	DBMaxIdleConnections int    `mapstructure:"DB_MAX_IDLE_CONNECTIONS"`
	DBHost               string `mapstructure:"DB_HOST"`
}

type PostgresConfig struct {
	postgres.PostgresConfig
}

func (cfg AppConfig) Init(filename string) error {
	viper.SetConfigName(filename)
	viper.AddConfigPath("./")
	viper.AddConfigPath("./../")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return nil
}

func (cfg AppConfig) Validate() error {
	return nil
}

func main() {
	fmt.Println("initiating the project")
	cfg := AppConfig{}
	err := cfg.Init("application")
	if err != nil {
		fmt.Println("viper init ", err.Error())
		return
	}
	err = config.Load(&cfg)
	if err != nil {
		fmt.Println("config load ", err.Error())
		return
	}

	fmt.Println(cfg)

	postgresConfig := postgres.GetPostgresConfig(cfg.DBPort, cfg.DBPool, cfg.DBMaxIdleConnections, cfg.DBUser, "", cfg.DBHost, cfg.DBName)

	conn, err := postgres.NewPostgres(postgresConfig)

	if err != nil {
		fmt.Println(err.Error())
	}
	err = conn.Ping()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("successfully create postgres connection")
	}
}
