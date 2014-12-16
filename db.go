package main

import (
	"dropler-new/store"
	"dropler-new/users"
)

func setupDb() {
	// Setup the db config options for each component
	users.SetupDb(store.Db)
	return
}
