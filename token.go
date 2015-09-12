package main

import "droppio/utils/time"

// Token struct
type Token struct {
	Id          int64  `json:"id"db:"Id"`
	ClientID    string `json:"client"db:"-"`
	Code        string `json:"access_code"`
	ExpiresIn   int32  `json:"expires"`
	Scope       string `json:"scope"`
	RedirectUri string `json:"redirect_uri"`
	State       string `json:"state"`
	models.TimeStamp
}

func (t *Token) Insert() error {
	t.UpdateTime()

	// run the DB insert function
	err := Db.Insert(t)
	if err != nil {
		return err
	}

	return nil
}

func (t *Token) GetByCode(code string) error {
	err := Db.SelectOne(t, "SELECT * FROM access_tokens WHERE code=$1", code)
	if err != nil {
		return err
	}

	return nil
}
