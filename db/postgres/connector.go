package db

import (
	"database/sql"
	"fmt"
	"log"
	"memorymaps-backend/config"

	// Ued for connecting to postgres server
	_ "github.com/lib/pq"
)

// GetDBConn : Used for DB connections
func GetDBConn(dbname string) (*sql.DB, error) {

	dbinfo := fmt.Sprintf("dbname=%s sslmode=disable", dbname)

	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		fmt.Println("Error in connecting to PostGres DB")
		return nil, err
	}

	// Open doesn't open a connection. Validate DSN data:
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil

}

// CreateTablesIfNotExists : Creates an index table if not present
func CreateTablesIfNotExists() error {

	log.Println("Validating the presence of tables on server ...")

	// Create DB conn
	db, err := GetDBConn(config.DBName)

	if err != nil {
		log.Println("Error Connecting to DB")
		log.Println(err)
		return err
	}

	// Defer db close
	defer db.Close()

	// Creating the tables
	// TextMemory Table
	_, err = db.Exec(
		"CREATE TABLE IF NOT EXISTS TextMemory(ID SERIAL PRIMARY KEY, TextMem VARCHAR(1000) NOT NULL, CreationTime timestamp default current_timestamp);")

	if err != nil {
		log.Println(err)
		return err
	}

	log.Println("Tables are ready as required.")

	return nil
}
