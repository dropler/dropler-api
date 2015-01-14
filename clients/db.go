package clients

import "dropler/store"

func init() {
	store.Db.AddTableWithName(Client{}, "clients").SetKeys(true, "Id")
	return
}
