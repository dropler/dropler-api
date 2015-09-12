package main

import (
	"errors"
	"fmt"

	"github.com/RangelReale/osin"
	"github.com/gin-gonic/gin"
)

var (
	// OauthServer public instance of an osin.Server
	OauthServer = newAuthServer()
)

func AccessToken(c *gin.Context) {
	var (
		oauth = OauthServer
		resp  = oauth.NewResponse()
	)
	defer resp.Close()

	if ar := oauth.HandleAccessRequest(resp, c.Request); ar != nil {
		switch ar.Type {
		case osin.AUTHORIZATION_CODE:
			ar.Authorized = true
		case osin.REFRESH_TOKEN:
			ar.Authorized = true
		case osin.PASSWORD:
			if authorizeUser(ar.Username, ar.Password) {
				ar.Authorized = true
			}
		case osin.CLIENT_CREDENTIALS:
			ar.Authorized = true
		}
		oauth.FinishAccessRequest(resp, c.Request, ar)
	}

	if resp.IsError && resp.InternalError != nil {
		fmt.Printf("ERROR: %s\n", resp.InternalError)
	}

	osin.OutputJSON(resp, c.Writer, c.Request)
}

func authorizeUser(username, password string) bool {
	var u User

	err := u.FindByEmail(username)
	if err != nil {
		return false
	}

	err = u.CheckPassword(password)
	if err != nil {
		return false
	}

	return true
}

func newAuthServer() *osin.Server {
	authConfig := osin.NewServerConfig()

	authConfig.AllowedAuthorizeTypes = osin.AllowedAuthorizeType{osin.CODE, osin.TOKEN}

	authConfig.AllowedAccessTypes = osin.AllowedAccessType{osin.AUTHORIZATION_CODE, osin.REFRESH_TOKEN, osin.PASSWORD}

	authConfig.AllowClientSecretInParams = true
	authConfig.AllowGetAccessRequest = true

	server := osin.NewServer(authConfig, newAuthStorage())
	return server
}

func newAuthStorage() *oauthStorage {
	s := &oauthStorage{}
	return s
}

type oauthStorage struct{}

func (s *oauthStorage) Clone() osin.Storage {
	return s
}

func (s *oauthStorage) Close() {
}

func (s *oauthStorage) GetClient(id string) (osin.Client, error) {
	fmt.Printf("GetClient: %s\n", id)

	c := Client{}

	err := c.GetClientByID(id)
	if err != nil {
		return nil, errors.New("Client not found")
	}

	fmt.Printf("ClientFound: %s\n", c.ClientSecret)

	return c, nil
}

func (s *oauthStorage) SetClient(id string, client osin.Client) error {
	fmt.Printf("SetClient: %s\n", id)
	return nil
}

func (s *oauthStorage) SaveAuthorize(data *osin.AuthorizeData) error {
	fmt.Printf("SaveAuthorize: %s\n", data.Code)
	return nil
}

func (s *oauthStorage) LoadAuthorize(code string) (*osin.AuthorizeData, error) {
	fmt.Printf("LoadAuthorize: %s\n", code)
	return nil, errors.New("Authorize not found")
}

func (s *oauthStorage) RemoveAuthorize(code string) error {
	fmt.Printf("RemoveAuthorize: %s\n", code)
	return nil
}

func (s *oauthStorage) SaveAccess(data *osin.AccessData) error {
	fmt.Printf("SaveAccess: %s\n", data.AccessToken)
	accessToken := Token{
		ClientID:    data.Client.GetId(),
		Code:        data.AccessToken,
		ExpiresIn:   data.ExpiresIn,
		Scope:       data.Scope,
		RedirectUri: data.RedirectUri,
	}

	err := accessToken.Insert()
	if err != nil {
		return err
	}

	return nil
}

func (s *oauthStorage) LoadAccess(code string) (*osin.AccessData, error) {
	fmt.Printf("LoadAccess: %s\n", code)
	return nil, errors.New("Access not found")
}

func (s *oauthStorage) RemoveAccess(code string) error {
	fmt.Printf("RemoveAccess: %s\n", code)
	return nil
}

func (s *oauthStorage) LoadRefresh(code string) (*osin.AccessData, error) {
	fmt.Printf("LoadRefresh: %s\n", code)
	return nil, errors.New("Refresh not found")
}

func (s *oauthStorage) RemoveRefresh(code string) error {
	fmt.Printf("RemoveRefresh: %s\n", code)
	return nil
}
