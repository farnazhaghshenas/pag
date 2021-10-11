package config

import "time"

type config struct {
	App App `mapstructure:"app" validate:"required"`
	DB  DB  `mapstructure:"DB" validate:"required"`
}

// GetAppConfig will return the nats configs
func (conf *config) GetAppConfig() App {
	return conf.App
}

// GetDBConfig will return the web engage configs
func (conf *config) GetDBConfig() DB {
	return conf.DB
}

type App struct {
	Port string `mapstructure:"port" validate:"required"`
}

type DB struct {
	Name               string        `mapstructure:"DATABASE"`
	Host               string        `mapstructure:"HOST"`
	Port               int           `mapstructure:"PORT"`
	Username           string        `mapstructure:"USERNAME"`
	Password           string        `mapstructure:"PASSWORD" default:""`
	MaxIdleConnections int           `mapstructure:"MAX_IDLE_CONNECTIONS"`
	MaxOpenConnections int           `mapstructure:"OPEN_CONNECTIONS"`
	ConnMaxLifeTime    time.Duration `mapstructure:"MAX_LIFE_TIME"`
	ReadTimeout        string        `mapstructure:"READ_TIMEOUT"`
}
