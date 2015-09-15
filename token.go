package main

import "droppio/utils/time"

// AccessToken struct
type AccessToken struct {
	Id        int64  `json:"id"db:"Id"`
	UserID    int64  `json:"user"`
	Token     string `json:"access_token"`
	ClientID  string `json:"client"`
	ExpiresIn int32  `json:"expires"`
	Scope     string `json:"scope"`
	models.TimeStamp
}

func (t *AccessToken) Insert() error {
	t.UpdateTime()

	// run the DB insert function
	err := Db.Insert(t)
	if err != nil {
		return err
	}

	return nil
}

func GetUserByAccessToken(code string) (*User, error) {
	t := &AccessToken{}
	err := Db.SelectOne(t, "SELECT * FROM access_tokens WHERE Token=$1", code)
	if err != nil {
		return nil, err
	}

	u := &User{}

	err = u.GetById(t.UserID)
	if err != nil {
		return nil, err
	}

	return u, nil
}
