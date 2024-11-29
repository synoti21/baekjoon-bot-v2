package config

type Database string

const (
	Mongo = Database("mongo")
	Local = Database("local")
)
