package db

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

//The DB struct contains a connect with postgres db or a error connection.
type DB struct {
	Connect *sqlx.DB
	Error   error
}

var db *DB
var once sync.Once

// GetDBConnect returns a db connection struct. If there was any error then a struct will have error.
func GetDBConnect() *DB {

	once.Do(func() {
		strConnect := fmt.Sprintf("host=db sslmode=disable user=%s password=%s dbname=%s", getEnv("POSTGRES_USER"), getEnv("POSTGRES_PASSWORD"), getEnv("POSTGRES_DB"))
		db = &DB{}
		db.Connect, db.Error = sqlx.Connect("postgres", strConnect)
	})

	return db
}

// getEnv return a value from env list by key.
func getEnv(key string) string {
	value := os.Getenv(key)

	if value == "" {
		log.Fatalf("%s environment variable not set.", key)
	}

	return value
}
