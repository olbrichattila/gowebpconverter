package dbstorage

import (
	"database/sql"
	"webpcdn/internal/ports"

	// this needs to be permanently imported but not referenced
	_ "github.com/mattn/go-sqlite3" // SQLite driver
)

const (
	insertSQL = "INSERT INTO cached_images (key, image) values (?,?)"
	selectSQL = "SELECT image FROM cached_images WHERE key = ?"
	existsSQL = "SELECT COUNT(id) FROM cached_images WHERE KEY = ?"
)

func New() ports.Storer {
	return &store{}
}

type store struct {
}

// isExists implements ports.Storer.
func (s *store) IsExists(filename string) bool {
	db, err := s.dbOpen()
	if err != nil {
		return false
	}
	defer db.Close()

	stmt, err := db.Prepare(existsSQL)
	if err != nil {
		return false
	}

	defer stmt.Close()

	var res bool

	row := stmt.QueryRow(filename)
	if err != nil {
		return false
	}

	row.Scan(&res)

	return res
}

// Read implements ports.Storer.
func (s *store) Read(filename string) ([]byte, error) {
	db, err := s.dbOpen()
	if err != nil {

		return nil, err
	}
	defer db.Close()

	stmt, err := db.Prepare(selectSQL)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	var res []byte
	row := stmt.QueryRow(filename)
	row.Scan(&res)

	return res, nil
}

// Write implements ports.Storer.
func (s *store) Write(filename string, data []byte) error {
	db, err := s.dbOpen()
	if err != nil {
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare(insertSQL)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(filename, data)
	if err != nil {
		return err
	}

	return nil
}

func (s *store) dbOpen() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./db/database.sqlite")
	if err != nil {
		return nil, err
	}

	return db, err
}
