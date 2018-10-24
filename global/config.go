package global

import (
	"github.com/avasapollo/beers-api/database"
	"github.com/kelseyhightower/envconfig"
)

type AppConfig struct {
	AppName string `envconfig:"APP_NAME" default:"beers-api"`
	Env     string `envconfig:"ENV" default:"dev"`
	Port    int    `envconfig:"PORT" default:"8080""`
	*database.MongoConfig
}

func (config AppConfig) MongoUrlIsSet() bool {
	if config.MongoConfig == nil {
		return false
	}
	if config.MongoUrl == "" {
		return false
	}
	return true
}

func NewAppConfig() (*AppConfig, error) {
	config := new(AppConfig)
	if err := envconfig.Process("", config); err != nil {
		return nil, err
	}
	return config, nil
}
