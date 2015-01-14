package oauth

import (
	"dropler/models/token"
	"dropler/store"
)

func init() {
	store.Db.AddTableWithName(models.Token{}, "access_tokens").SetKeys(true, "Id")
	return
}
