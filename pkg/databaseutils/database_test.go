package databaseutils

import (
	"log"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

const databaseConfigurationFile = "./sample_config.json"

func TestLoadConfigurationFile(t *testing.T) {
	var database Database
	err := database.LoadConfigurationFile(databaseConfigurationFile)
	if err != nil {
		t.Error(err)
		return
	}

	if database.DatabaseConfiguration == nil {
		t.Error("config is nil")
		return
	}

	if database.DatabaseConfiguration.Connection.MaxIdleConns != 0 {
		t.Error("config.Connection.MaxIddleConnections != 0")
		return
	}

	log.Println("Configuration: ", database.DatabaseConfiguration)

	log.Println("Connection string: " + database.DatabaseConfiguration.GetConnectionString())
}

func TestDatabaseConnection(t *testing.T) {
	var database Database
	err := database.LoadConfigurationFile(databaseConfigurationFile)
	if err != nil {
		t.Error(err)
		return
	}

	err = database.Connect()
	if err != nil {
		t.Error(err)
		return
	}

	if database.dataBaseHandler == nil {
		t.Error("database.DataBaseHandler is nil")
		return
	}

	log.Println("Database connection established.")

	err = database.dataBaseHandler.Close()
	if err != nil {
		t.Error("Error closing database connection")
		return
	}

	log.Println("Database disconnected.")
}
