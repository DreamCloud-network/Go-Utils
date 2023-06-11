package databaseutils

import (
	"database/sql"
)

// Database - Struct repreenting a database connection
type Database struct {
	DatabaseConfiguration *DatabaseConfiguration
	dataBaseHandler       *sql.DB
}

// DatabaseConfiguration - Struct with all database configuration information
type DatabaseConfiguration struct {
	Connection ConnectionConfig `json:"connection,omitempty"`
	Access     AccessConfig     `json:"access,omitempty"`
	Server     ServerConfig     `json:"server,omitempty"`
}

// ConnectionConfig - Struct with connection configuration information
type ConnectionConfig struct {
	ConnMaxIdleTime int `json:"connMaxIdleTime,omitempty"`
	ConnMaxLifetime int `json:"connMaxLifetime,omitempty"`
	MaxIdleConns    int `json:"maxIdleConns,omitempty"`
	MaxOpenConns    int `json:"maxOpenConns,omitempty"`
}

// AccerssConfig - Struct with database access configuration information
type AccessConfig struct {
	User     string `json:"user,omitempty"`
	Password string `json:"password,omitempty"`
	DbName   string `json:"dbname,omitempty"`
}

// ServerConfig - Struct with server address configuration information
type ServerConfig struct {
	Address string `json:"address,omitempty"`
	Port    int    `json:"port,omitempty"`
	DBType  string `json:"dbtype,omitempty"`
}
