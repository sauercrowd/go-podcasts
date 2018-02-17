package storage

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/sauercrowd/go-podcasts/pkg/flags"
)

func Setup(f *flags.Vars) (error, *sql.DB) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", f.PostgresUser, f.PostgresPassword, f.PostgresHost, f.PostgresPort, f.PostgresUser)
	log.Println(connStr)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return err, nil
	}
	if err := createDatabaseIfNotExist(db, f.PostgresDatabase); err != nil {
		return err, nil
	}
	if err := db.Close(); err != nil {
		return err, nil
	}
	connStr = fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", f.PostgresUser, f.PostgresPassword, f.PostgresHost, f.PostgresPort, f.PostgresDatabase)
	db, err = sql.Open("postgres", connStr)
	if err := createTablesIfNotExist(db); err != nil {
		return err, nil
	}
	return nil, db
}

func createDatabaseIfNotExist(db *sql.DB, name string) error {
	var count int64
	err := db.QueryRow("SELECT COUNT(1) FROM pg_database WHERE datname = $1", name).Scan(&count)
	//return if database exists or error happend
	if err != nil || count == 1 {
		if err == nil {
			err = db.Close()
		}
		return err
	}
	err = db.QueryRow(fmt.Sprintf("CREATE DATABASE %s", name)).Scan()
	if err != nil && err != sql.ErrNoRows {
		return err
	}
	return nil
}

func createTablesIfNotExist(db *sql.DB) error {
	err := db.QueryRow("CREATE TABLE IF NOT EXISTS podcasts(podcastid serial, podcastlink text PRIMARY KEY, podcastname text, lang text, description text, imageurl text, updated timestamp)").Scan()
	if err != nil && err != sql.ErrNoRows {
		return err
	}
	err = db.QueryRow("CREATE TABLE IF NOT EXISTS episodes( audiourl text PRIMARY KEY, episodelink text, podcastlink text references podcasts(podcastlink) NOT NULL, episodename text, episodedescription text, pubdate timestamp)").Scan()
	if err != nil && err != sql.ErrNoRows {
		return err
	}
	return nil
}
