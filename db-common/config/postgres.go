package config

import "fmt"

type PostgresConfig struct {
	Host     string
	User     string
	Password string
	Name     string
	Port     string
}

func (conf *PostgresConfig) ToConnectionString() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		conf.Host,
		conf.User,
		conf.Password,
		conf.Name,
		conf.Port,
	)
}
