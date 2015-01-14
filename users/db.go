package users

import "dropler/store"

func init() {
	store.Db.AddTableWithName(User{}, "users").SetKeys(true, "Id")
	return
}
