package config

type Conf interface {
	GetAppConfig() App
	GetDBConfig() DB
}
