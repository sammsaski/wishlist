/**
* This file is responsible for creating an instance of our database
* and establishing a connection with it.
*/

package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	// The host name is no longer "localhost" when our Postgres instance is containerized.
	// Instead, the host name becomes the name of our service as defined in our docker-compose.
	// Hence, if you check the docker-compose, our DB service is named "database".
	host = "database"
	port = 5432
)

// returned when we request a row that doesn't exist
var ErrNoMatch = fmt.Errorf("no matching record")

/**
*	Holds the database instance.
*/
type Database struct {
	Conn *sql.DB
}

/**
*	Establish a connection with the database.
*/
func Initialize(username, password, database string) (Database, error) {
	db := Database{}
	dbInfo := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable", username, password, host, port, database)
	fmt.Println(dbInfo)
	conn, err := sql.Open("postgres", dbInfo)
	fmt.Println("Connection open")
	if err != nil {
		return db, err
	}
	db.Conn = conn
	err = db.Conn.Ping()
	if err != nil {
		return db, err
	}
	fmt.Println("Database connection established")
	return db, nil
}
