package oauth

import (
	"dropler/store"
	"dropler/token"
)

func init() {
	store.Db.AddTableWithName(models.Token{}, "access_tokens").SetKeys(true, "Id")
	return
}
