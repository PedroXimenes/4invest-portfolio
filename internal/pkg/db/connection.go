package db

import (
	"database/sql"

	log "github.com/sirupsen/logrus"

	"cloud.google.com/go/cloudsqlconn"
	"cloud.google.com/go/cloudsqlconn/postgres/pgxv4"
)

func init() {
	cleanup, err := pgxv4.RegisterDriver("cloudsql-postgres", cloudsqlconn.WithIAMAuthN())
	defer cleanup()
	if err != nil {
		log.Fatal(err)
	}
}
func OpenConnection() (*sql.DB, error) {
	db, err := sql.Open(
		"cloudsql-postgres",
		"host=sincere-almanac-401411:us-central1:four-invest user=postgres password=m-,b(<(vIgsuXCV< dbname=four-invest sslmode=disable",
	)
	if err != nil {
		return nil, err
	}
	return db, nil
}
