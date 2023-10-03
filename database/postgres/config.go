package postgres

import (
	"fmt"
)

type PostgresConfig struct {
	Host               string
	Port               int
	Name               string
	Username           string
	Password           string
	MaxPoolSize        int
	MaxIdleConnections int
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
