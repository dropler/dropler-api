package drops

import (
	"github.com/begizi/gorp"
)

func SetupDb(g *gorp.DbMap) {
	g.AddTableWithName(Drop{}, "drops").SetKeys(true, "Id")
	return
}
