package config

import (
	"context"

	"github.com/kelseyhightower/envconfig"
)

type DBConfiguration struct {
	Endpoint string `envconfig:"MONGO_DB_ENDPOINT" required:"true"`
}

func GetDBConfig(ctx context.Context) DBConfiguration {
	dbConfig := DBConfiguration{}
	err := envconfig.Process("", &dbConfig)
	if err != nil {
		panic(err)
	}
	return dbConfig
}
