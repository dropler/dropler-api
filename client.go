package main

import (
	"droppio/utils/time"
	"encoding/base64"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/pborman/uuid"
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
	_, err := Db.Select(c, "SELECT * FROM CLIENTS ORDER BY CreatedAt DESC")
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
	err := Db.Insert(c)

	return err
}

func (c *Client) GetById(id string) error {
	err := Db.SelectOne(c, "SELECT * FROM clients WHERE ID=$1", id)
	return err
}

// GetClientByID Method for returning a single client row.
func (c *Client) GetClientByID(id string) error {

	err := Db.SelectOne(c, "SELECT * FROM clients WHERE ClientID=$1", id)

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

// List Method for returning all clients as array of objects
func ListClient(c *gin.Context) {
	d := ClientList{}

	err := d.List()
	if err != nil {
		c.JSON(500, gin.H{"error": "Problem fetching clients. No database connection?"})
		return
	}

	c.JSON(200, d)
	return
}

// Create Method for creating a new client
func CreateClient(c *gin.Context) {
	d := Client{}

	// Bind the incoming form data with the client model
	c.BindWith(&d, binding.Form)

	err := d.Create()
	if err != nil {
		panic(err)
	}

	// Return the client model as json to the client
	c.JSON(200, d)
	return
}

func GetClient(c *gin.Context) {
	d := Client{}
	id := c.Params.ByName("id")

	err := d.GetById(id)
	if err != nil {
		c.JSON(500, gin.H{"error": "Problem fetching client. No database connection?"})
		return
	}

	c.JSON(200, d)
}
