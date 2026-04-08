package db

import (
	"database/sql"
	"os"

	_ "modernc.org/sqlite"
)

const schema string = `CREATE TABLE scheduler (
	id INTEGER PRIMARY KEY AUTOINCREMENT, 
	date CHAR(8) NOT NULL DEFAULT "",
	title VARCHAR(128) NOT NULL,
	comment TEXT NOT NULL DEFAULT "",
	repeat VARCHAR(128) NOT NULL DEFAULT ""
);

CREATE INDEX date_idx ON scheduler(date);`

var db *sql.DB

func Init(dbFile string) error {
	_, err := os.Stat(dbFile)
	var install bool
	if err != nil {
		install = true
	}

	db, err = sql.Open("sqlite", dbFile)
	if err != nil {
		return err
	}

	if install {
		if _, err := db.Exec(schema); err != nil {
			return err
		}
	}

	return nil
}

func Close() error {
	return db.Close()
}
