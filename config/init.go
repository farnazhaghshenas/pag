package config

import (
	"strings"

	"github.com/spf13/viper"
)

// NewViperConfiguration is the Viper Configuration factory method.
func NewViperConfiguration() (*config, error) {
	configuration := new(config)

	vip := viper.New()

	//setDefaultValues(vip)

	//vip.AutomaticEnv()
	vip.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	vip.SetConfigName("config")
	vip.AddConfigPath(".")

	err := vip.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = vip.Unmarshal(configuration)
	if err != nil {
		return nil, err
	}

	return configuration, nil
}

//func setDefaultValues(v *viper.Viper) {
//	v.SetDefault("app.port", "8080")
//
//	v.SetDefault("db.DATABASE", "test")
//	v.SetDefault("db.HOST", "localhost")
//	v.SetDefault("db.PORT", 3306)
//	v.SetDefault("db.USERNAME", "root")
//	v.SetDefault("db.PASSWORD", "")
//	v.SetDefault("db.MAX_IDLE_CONNECTIONS", 10)
//	v.SetDefault("db.OPEN_CONNECTIONS", 0)
//	v.SetDefault("db.MAX_LIFE_TIME", 0)
//	v.SetDefault("db.READ_TIMEOUT", "5s")
//}
