package oauth

import (
	"dropler/clients"
	"errors"
	"fmt"
	"github.com/RangelReale/osin"
)

var (
	// OauthServer public instance of an osin.Server
	OauthServer = newAuthServer()
)

type oauthStorage struct{}

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

func (s *oauthStorage) Clone() osin.Storage {
	return s
}

func (s *oauthStorage) Close() {
}

func (s *oauthStorage) GetClient(id string) (osin.Client, error) {
	fmt.Printf("GetClient: %s\n", id)

	c := clients.Client{}

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
