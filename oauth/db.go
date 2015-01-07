package oauth

import (
	"dropler/models/token"
	"github.com/begizi/gorp"
)

func SetupDb(g *gorp.DbMap) {
	g.AddTableWithName(models.Token{}, "access_tokens").SetKeys(true, "Id")
	return
}
