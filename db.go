package main

import (
	"dropler-new/clients"
	"dropler-new/drops"
	"dropler-new/store"
	"dropler-new/users"
)

func setupDb() {
	// Setup the db config options for each component
	clients.SetupDb(store.Db)
	drops.SetupDb(store.Db)
	users.SetupDb(store.Db)
	return
}
