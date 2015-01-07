package models

import (
	"dropler/clients"
	"dropler/models/time"
)

// Token struct
type Token struct {
	Id          int64          `json:"id"db:"Id"`
	Client      clients.Client `json:"client"db:"-"`
	Code        string         `json:"access_code"`
	ExpiresIn   int32          `json:"expires"`
	Scope       string         `json:"scope"`
	RedirectUri string         `json:"redirect_uri"`
	State       string         `json:"state"`
	models.TimeStamp
}
