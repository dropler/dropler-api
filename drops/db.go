package drops

import "dropler/store"

func init() {
	store.Db.AddTableWithName(Drop{}, "drops").SetKeys(true, "Id")
	return
}
