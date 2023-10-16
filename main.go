package main

import (
	"fmt"
	"github.com/modern-era-devs/go-utils/config"
	"github.com/modern-era-devs/go-utils/database/postgres"
	"github.com/spf13/viper"
	//_ "gopkg.in/yaml.v2"
)

type AppConfig struct {
	PostgresCfg postgres.PostgresConfig `mapstructure:"POSTGRES"`
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

	//postgresConfig := postgres.GetPostgresConfig(cfg.PostgresCfg.Port, cfg.PostgresCfg.MaxPoolSize, cfg.PostgresCfg.MaxIdleConnections, cfg.PostgresCfg.Username, cfg.PostgresCfg.Password, cfg.PostgresCfg.Host, cfg.PostgresCfg.Name)

	//conn, err := postgres.NewPostgres(postgresConfig)

	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	//err = conn.Ping()
	//if err != nil {
	//	fmt.Println(err.Error())
	//} else {
	//	fmt.Println("successfully create postgres connection")
	//}
}
