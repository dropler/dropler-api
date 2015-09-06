package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/begizi/gorp"

	// Database needed in sub-package, store
	_ "github.com/lib/pq"
)

var (
	// Db is a global variable to access the db context
	Db = NewDbContext()
)

func init() {
	Db.AddTableWithName(User{}, "users").SetKeys(true, "Id")
	Db.AddTableWithName(Token{}, "access_tokens").SetKeys(true, "Id")
	Db.AddTableWithName(Drop{}, "drops").SetKeys(true, "Id")
	Db.AddTableWithName(Client{}, "clients").SetKeys(true, "Id")
}

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
