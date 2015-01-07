package main

import (
	"dropler/clients"
	"dropler/drops"
	"dropler/oauth"
	"dropler/store"
	"dropler/users"
)

func setupDb() {
	// Setup the db config options for each component
	clients.SetupDb(store.Db)
	drops.SetupDb(store.Db)
	oauth.SetupDb(store.Db)
	users.SetupDb(store.Db)
	return
}
