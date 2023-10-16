package postgres

import (
	"fmt"
	//_ "gopkg.in/yaml.v2"
)

type PostgresConfig struct {
	Host               string `mapstructure:"HOST"`
	Port               int    `mapstructure:"PORT"`
	Name               string `mapstructure:"NAME"`
	Username           string `mapstructure:"USERNAME"`
	Password           string `mapstructure:"PASSWORD"`
	MaxPoolSize        int    `mapstructure:"MAX_POOL_SIZE"`
	MaxIdleConnections int    `mapstructure:"MAX_IDLE_CONNECTIONS"`
}

func GetPostgresConfig(port, poolSize, maxIdleConn int, user, password, host, dbName string) PostgresConfig {
	return PostgresConfig{
		Host:               host,
		Port:               port,
		Name:               dbName,
		Username:           user,
		Password:           password,
		MaxPoolSize:        poolSize,
		MaxIdleConnections: maxIdleConn,
	}
}

func (cfg PostgresConfig) GetHost() string {
	return cfg.Host
}

func (cfg PostgresConfig) GetPort() int {
	return cfg.Port
}

func (cfg PostgresConfig) GetName() string {
	return cfg.Name
}

func (cfg PostgresConfig) GetMaxPoolSize() int {
	return cfg.MaxPoolSize
}

func (cfg PostgresConfig) GetMaxIdleConnections() int {
	return cfg.MaxIdleConnections
}

func (cfg PostgresConfig) GetConnectionString() string {
	return fmt.Sprintf("dbname=%s user=%s password=%s host=%s sslmode=disable", cfg.Name, cfg.Username, cfg.Password, cfg.Host)
}

func (cfg PostgresConfig) GetConnectionURL() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Name)
}
