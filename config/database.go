package config

import (
	"os"
	"strconv"

	"github.com/synoti21/baekjoon-slack-bot/common/errors"
)

type DatabaseType string

const (
	DatabaseTypePostgres DatabaseType = "postgres"
	DatabaseTypeMySQL    DatabaseType = "mysql"
	DatabaseTypeMongoDB  DatabaseType = "mongodb"
	DatabaseTypeDryRun   DatabaseType = "dryrun"
)

type DatabaseClientConfig struct {
	Type     DatabaseType // Type of database (e.g., postgres, mysql, mongodb)
	Host     string       // Database server hostname or IP
	Port     int          // Port on which the database is running
	User     string       // Username for database authentication
	Password string       // Password for database authentication
	DBName   string       // Name of the specific database to use
	SSLMode  string       // SSL mode (e.g., "disable", "require")
}

func NewDatabaseClientConfig() *DatabaseClientConfig {
	dbtype := databaseType()
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	portNum, err := strconv.ParseInt(port, 0, 64)
	if err != nil {
		panic("Invalid DB Port")
	}
	user := os.Getenv("DB_USER")
	pwd := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	sslmode := os.Getenv("DB_SSL_MODE")

	return &DatabaseClientConfig{
		Type:     dbtype,
		Host:     host,
		Port:     int(portNum),
		User:     user,
		Password: pwd,
		DBName:   dbname,
		SSLMode:  sslmode,
	}

}

// Validate checks the configuration fields for validity
func (c *DatabaseClientConfig) Validate() error {
	if c.Type == "" {
		return errors.NewInternalServerError("database type cannot be empty")
	}
	switch c.Type {
	case DatabaseTypePostgres, DatabaseTypeMySQL:
		if c.Host == "" {
			return errors.NewInternalServerError("host cannot be empty")
		}
		if c.Port <= 0 || c.Port > 65535 {
			return errors.NewInternalServerError("port must be a valid number between 1 and 65535")
		}
		if c.User == "" {
			return errors.NewInternalServerError("user cannot be empty")
		}
		if c.DBName == "" {
			return errors.NewInternalServerError("database name cannot be empty")
		}
		if c.SSLMode == "" {
			return errors.NewInternalServerError("SSL mode cannot be empty")
		}
	case DatabaseTypeMongoDB:
		if c.Host == "" {
			return errors.NewInternalServerError("host cannot be empty")
		}
		if c.Port <= 0 || c.Port > 65535 {
			return errors.NewInternalServerError("port must be a valid number between 1 and 65535")
		}
		if c.DBName == "" {
			return errors.NewInternalServerError("database name cannot be empty")
		}
	default:
		return errors.NewInternalServerError("unsupported database type")
	}
	return nil
}

func databaseType() DatabaseType {
	db := os.Getenv("DATABASE_MODE")
	if db == "" {
		panic("DATABASE_MODE not set")
	}
	switch db {
	case "mongo":
		return DatabaseTypeMongoDB
	case "mysql":
		return DatabaseTypeMySQL
	case "postgres":
		return DatabaseTypePostgres
	case "local":
		return DatabaseTypeDryRun
	default:
		panic("Invalid database mode")
	}
}
