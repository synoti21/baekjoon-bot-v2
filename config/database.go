package config

import "os"

type Database string

const (
	Mongo = Database("mongo")
	Local = Database("local")
)

func (c *HandlerConfig) Database() Database {
	db := os.Getenv("DATABASE_MODE")
	if db == "" {
		panic("DATABASE_MODE not set")
	}
	switch db {
	case "MONGO":
		return Mongo
	case "LOCAL":
		return Local
	default:
		panic("Invalid database mode")
	}
}
