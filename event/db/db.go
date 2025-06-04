package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error

	// Open (or create) the SQLite database file “api.db”
	DB, err = sql.Open("sqlite3", "api.db")
	if err != nil {
		panic("Could not connect to database: " + err.Error())
	}

	// Make sure SQLite enforces foreign keys
	// This must be run before creating tables or inserting rows that use FKs.
	if _, err := DB.Exec("PRAGMA foreign_keys = ON;"); err != nil {
		panic("Failed to enable foreign keys: " + err.Error())
	}

	// Pool settings (optional but recommended)
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	// Create both tables
	createTables()
}

func createTables() {
	// 1) Create the users table
	createUserTable := `
    CREATE TABLE IF NOT EXISTS users (
        id       INTEGER PRIMARY KEY AUTOINCREMENT,
        email    TEXT    NOT NULL UNIQUE,
        password TEXT    NOT NULL
    );
    `

	// 2) Create the events table with a foreign key to users(id)
	createEventsTable := `
    CREATE TABLE IF NOT EXISTS events (
        id          INTEGER  PRIMARY KEY AUTOINCREMENT,
        name        TEXT     NOT NULL,
        description TEXT     NOT NULL,
        location    TEXT     NOT NULL,
        dateTime    DATETIME NOT NULL,
        user_id     INTEGER,
        FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE
    );
    `

	createRegistrationsTable := `
		CREATE TABLE IF NOT EXISTS registrations (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    event_id INTEGER,
    user_id INTEGER,
    FOREIGN KEY(event_id) REFERENCES events(id),
    FOREIGN KEY(user_id) REFERENCES users(id)
)
`

	// Execute creation of users table
	if _, err := DB.Exec(createUserTable); err != nil {
		panic(fmt.Sprintf("Could not create users table: %v", err))
	}

	// Execute creation of events table
	if _, err := DB.Exec(createEventsTable); err != nil {
		panic(fmt.Sprintf("Could not create events table: %v", err))
	}

	_, err := DB.Exec(createRegistrationsTable)
	if err != nil {
		panic("Could not create registrations table.")
	}
}
