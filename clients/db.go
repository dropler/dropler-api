package clients

import (
	"github.com/begizi/gorp"
)

func SetupDb(g *gorp.DbMap) {
	g.AddTableWithName(Client{}, "clients").SetKeys(true, "Id")
	return
}
