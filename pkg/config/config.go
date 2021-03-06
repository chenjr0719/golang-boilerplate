package config

import (
	"github.com/spf13/viper"
)

type config struct {
	DatabaseURI string `mapstructure:"DATABASE_URI"`
	Mode        string `mapstructure:"MODE"`
	LogLevel    string `mapstructure:"LOG_lEVEL"`
}

var Config config = LoadConfig()

func newViper() *viper.Viper {
	v := viper.New()
	v.AddConfigPath(".")
	v.SetConfigName(".env")
	v.SetConfigType("env")
	v.AutomaticEnv()

	setDefaults(v)

	return v
}

func setDefaults(v *viper.Viper) {
	v.SetDefault("DATABASE_URI", "sqlite://file::memory:?cache=shared")
	v.SetDefault("MODE", "release")
	v.SetDefault("LOG_lEVEL", "info")
}

func LoadConfig() config {
	v := newViper()

	config := config{}
	// Read from env first
	err := v.Unmarshal(&config)
	if err != nil {
		panic(err)
	}
	// Overwrite config by .env file
	v.ReadInConfig()

	return config
}
