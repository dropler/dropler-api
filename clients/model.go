package clients

import (
	"code.google.com/p/go-uuid/uuid"
	"dropler-new/models"
	"dropler-new/store"
	"encoding/base64"
	"log"
)

// Client model struct
type Client struct {
	ID           int64       `json:"id"db:"Id"`
	Name         string      `form:"name"json:"name"`
	RedirectURI  string      `form:"redirect_uri"json:"redirect_uri"`
	ClientID     string      `form:"-"json:"client_id"`
	ClientSecret string      `form:"-"json:"client_secret"`
	User         interface{} `form:"-"json:"-"db:"-"`
	models.TimeStamp
}

type ClientList []Client

// Fetch Method to return all clients from the db
func (c *ClientList) List() error {
	_, err := store.Db.Select(c, "SELECT * FROM CLIENTS ORDER BY CreatedAt DESC")
	if err != nil {
		log.Fatal(err)
	}
	return err
}

// Create Method to create a new client from the models Client struct
func (c *Client) Create() error {

	// run the UpdateTime ethod on the client model
	c.UpdateTime()

	// run the id generation during the creation of the client
	c.GenerateClientID()

	// run the secret generation during the creation of the client
	c.GenerateClientSecret()

	// run the DB insert function
	err := store.Db.Insert(c)

	return err
}

// GetClientByID Method for returning a single client row.
func (c *Client) GetClientByID(id string) error {

	err := store.Db.SelectOne(c, "SELECT * FROM clients WHERE ClientID=$1", id)

	return err
}

// GenerateClientID builds the clients secret
func (c *Client) GenerateClientID() {
	clientID := uuid.New()
	clientID = base64.StdEncoding.EncodeToString([]byte(clientID))
	clientID = clientID[:20]
	c.ClientID = clientID
}

// GenerateClientSecret builds the clients secret
func (c *Client) GenerateClientSecret() {
	accesstoken := uuid.New()
	accesstoken = base64.StdEncoding.EncodeToString([]byte(accesstoken))
	c.ClientSecret = accesstoken
}

func (c Client) GetId() string {
	return c.ClientID
}

func (c Client) GetRedirectUri() string {
	return c.RedirectURI
}

func (c Client) GetSecret() string {
	return c.ClientSecret
}

func (c Client) GetUserData() interface{} {
	return c.User
}
