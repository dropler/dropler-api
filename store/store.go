package store

import (
	"database/sql"
	"github.com/begizi/gorp"
	"log"
	"os"

	// Database needed in sub-package, store
	_ "github.com/lib/pq"
)

var (
	// Db is a global variable to access the db context
	Db = NewDbContext()
)

// NewDbContext establishes the database configuration
func NewDbContext() *gorp.DbMap {
	db, err := sql.Open("postgres", "dbname=dropler_dev sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}

	if os.Getenv("DEBUG") != "" {
		dbmap.TraceOn("[gorp]", log.New(os.Stdout, "[DATABASE] ", log.Lmicroseconds))
	}

	return dbmap
}
