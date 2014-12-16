package users

import (
	"github.com/begizi/gorp"
)

func SetupDb(g *gorp.DbMap) {
	g.AddTableWithName(User{}, "users").SetKeys(true, "Id")
	return
}
