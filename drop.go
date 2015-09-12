package main

import (
	"droppio/utils/time"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type Drop struct {
	ID        int64   `json:"id"db:"Id"`
	Name      string  `form:"name"json:"name"`
	Latitude  float32 `form:"lat"json:"lat"`
	Longitude float32 `form:"long"json:"long"`
	Radius    int     `form:"radius"json:"radius"`
	DropGeom  string  `form:"-"json:"-"`
	models.TimeStamp
}

type DropList []Drop

func (d *DropList) List() error {
	_, err := Db.Select(d, "SELECT * FROM drops ORDER BY CreatedAt DESC")
	if err != nil {
		log.Fatal(err)
	}
	return err
}

// Insert Method to create a new user from the models User struct
func (d *Drop) Insert() error {

	// run the UpdateTime ethod on the user model
	d.UpdateTime()

	// run the DB insert function
	err := Db.Insert(d)
	if err != nil {
		return err
	}

	return nil
}

func (d *Drop) GetById(id string) error {
	err := Db.SelectOne(d, "SELECT * FROM drops WHERE id=$1", id)
	if err != nil {
		return err
	}

	return nil
}

// List Method for returning all users as array of objects
func ListDrop(c *gin.Context) {
	d := DropList{}

	err := d.List()
	if err != nil {
		c.JSON(500, gin.H{"error": "Problem fetching drops. No database connection?"})
	}

	c.JSON(200, d)
}

func CreateDrop(c *gin.Context) {
	d := Drop{}

	c.BindWith(&d, binding.Form)

	err := d.Insert()
	if err != nil {
		c.JSON(500, gin.H{"error": "Problem fetching drops. No database connection?"})
		return
	}

	c.JSON(200, d)
}

func GetDrop(c *gin.Context) {
	d := Drop{}
	id := c.Params.ByName("id")

	err := d.GetById(id)
	if err != nil {
		c.JSON(500, gin.H{"error": "Problem fetching drops. No database connection?"})
		return
	}

	c.JSON(200, d)
}
