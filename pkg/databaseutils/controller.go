package databaseutils

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"reflect"
	"strconv"
	"time"

	"github.com/GreenMan-Network/Go-Utils/pkg/fileutils"
)

func (database *Database) LoadConfigurationFile(filePath string) error {
	// Open our jsonFile
	config, err := fileutils.ReadJsonFile(filePath)
	if err != nil {
		log.Println("databaseutils.LoadConfigurationFile - Error loading configuration file.")
		return err
	}

	if config == nil {
		log.Println("databaseutils.LoadConfigurationFile - Configuration file is nil.")
		return errors.New("configuration file is nil")
	}

	bs, _ := json.Marshal(config)

	var configStruct DatabaseConfiguration
	err = json.Unmarshal(bs, &configStruct)
	if err != nil {
		log.Println("databaseutils.LoadConfigurationFile - Error creating databaseconfig struct.")
		return err
	}

	/*
		TO DO: I don´t know why this is not working

		configStruct, ok := config.(DatabaseConfiguration)
		if !ok {
			log.Println("database.LoadConfiguration - Error converting configuration file to DatabaseConfiguration struct.")
			return errors.New("error converting configuration file to DatabaseConfiguration struct")
		}
	*/
	database.DatabaseConfiguration = &configStruct

	return nil
}

// GetConnectionString returns the connection string for the database
func (configuration *DatabaseConfiguration) GetConnectionString() string {
	return configuration.Access.User + ":" + configuration.Access.Password + "@tcp(" + configuration.Server.Address + ":" + strconv.Itoa(configuration.Server.Port) + ")/" + configuration.Access.DbName
}

// Connect opens and test a database conneciton.
func (database *Database) Connect() error {
	// Check if the configuration is loaded
	if database.DatabaseConfiguration == nil {
		log.Println("database.Connect - database.DatabaseConfiguration is nil.")
		return errors.New("database.DatabaseConfiguration is nil")
	}

	// Get the database handler
	db, err := sql.Open(database.DatabaseConfiguration.Server.DBType, database.DatabaseConfiguration.GetConnectionString())
	if err != nil {
		log.Println("database.Connect - Error connecting to the database.")
		return err
	}

	// Set the handler configuration
	db.SetConnMaxIdleTime(time.Duration(database.DatabaseConfiguration.Connection.ConnMaxIdleTime) * time.Second)
	db.SetConnMaxLifetime(time.Duration(database.DatabaseConfiguration.Connection.ConnMaxLifetime) * time.Second)

	db.SetMaxIdleConns(database.DatabaseConfiguration.Connection.MaxIdleConns)
	db.SetMaxOpenConns(database.DatabaseConfiguration.Connection.MaxOpenConns)

	// Test the connection
	err = db.Ping()
	if err != nil {
		log.Println("database.Connect - Error connecting to the database.")
		return err
	}

	// Set the database handler
	database.dataBaseHandler = db
	return nil
}

// Drop the table if it exists
func (database *Database) DropTable(table interface{}) error {
	statement := "DROP TABLE IF EXISTS " + reflect.TypeOf(table).Elem().Name()

	_, err := database.dataBaseHandler.Exec(statement)
	if err != nil {
		log.Println("database.DropTable - Error dropping table.")
		return err
	}

	return nil
}

/*
CREATE TABLE Persons (
    PersonID int,
    LastName varchar(255),
    FirstName varchar(255),
    Address varchar(255),
    City varchar(255)
);
*/

func (database *Database) CreateTable(table interface{}) error {
	statement := "CREATE TABLE " + reflect.TypeOf(table).Elem().Name() + " ("

	// Get the table fields
	fields := reflect.ValueOf(table).Elem()

	for i := 0; i < fields.NumField(); i++ {
		if i > 0 {
			statement += ","
		}
		statement += fields.Type().Field(i).Name + " " + fields.Type().Field(i).Type.String()
	}

	statement += ");"

	log.Println(statement)

	return nil
}
