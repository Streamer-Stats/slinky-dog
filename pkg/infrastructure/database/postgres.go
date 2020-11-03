package database

import (
	"github.com/go-pg/pg/v10"
)

// Database is my database struct
type Database struct {
	Username     string `json:"username,omitempty"`
	Password     string `json:"password,omitempty"`
	Name         string `json:"name,omitempty"`
	Host         string `json:"host,omitempty"`
	DatabaseType string `json:"databaseType,omitempty"`
}

func (database *Database) formatSQLConnection() *pg.Options {
	return &pg.Options{
		Addr:     "0.0.0.0:5432",
		User:     "postgres",
		Password: "123",
		Database: "streamerStats",
	}
}

// GetConnection get a database connection
func (database *Database) GetConnection() *pg.DB {
	return pg.Connect(database.formatSQLConnection())

}

// NewDatabase IoC
func NewDatabase() *Database {
	return &Database{}
}
