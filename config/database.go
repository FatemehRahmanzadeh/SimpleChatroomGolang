package config

import (
	"database/sql"
	"log"

	"github.com/FatemehRahmanzadeh/chat_sample/auth"
	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

func InitDB() *sql.DB {
	db, err := sql.Open("sqlite3", "./userdb.db")
	if err != nil {
		log.Fatal(err)
	}

	sqlStmt := `	
	CREATE TABLE IF NOT EXISTS user (
		id VARCHAR(255) NOT NULL PRIMARY KEY,
		username VARCHAR(255)  NULL UNIQUE,
		password VARCHARR(255)  NULL
	);
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
	}

	password, _ := auth.GeneratePassword("ftm1234")

	// insert data to database
	sqlStmt = `INSERT into user (id, username, password) VALUES
					('` + uuid.New().String() + `', 'zahra','` + password + `'),
					('` + uuid.New().String() + `', 'ali','` + password + `'),
					('` + uuid.New().String() + `', 'reza','` + password + `'),
					('` + uuid.New().String() + `', 'sara','` + password + `')`

	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
	}

	return db
}