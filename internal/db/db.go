package db

import (
	"database/sql"
	"log"
)

func InitDB(db *sql.DB) {
	initQuery := `PRAGMA foreign_keys = ON;
	PRAGMA journal_mode = WAL;
	PRAGMA synchronous = NORMAL;
	PRAGMA temp_store = MEMORY;
	`
	_, err := db.Exec(initQuery)
	if err != nil {
		log.Fatal(err)
	}
}

func CreateTable(db *sql.DB) {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS urls (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		url TEXT NOT NULL,
		short TEXT NOT NULL
	)`)
	if err != nil {
		log.Fatal(err)
	}
}

func AddURL(db *sql.DB, url string, short string) error {
	stmt, err := db.Prepare("INSERT INTO urls (url, short) VALUES (?, ?)")
	if err != nil {
		log.Println(err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(url, short)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func GetURL(db *sql.DB, short string) (string, error) {
	row := db.QueryRow("SELECT url FROM urls WHERE short = ?", short)
	var url string
	err := row.Scan(&url)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", nil
		}
		return "", err
	}
	return url, nil
}
