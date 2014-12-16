package drops

import (
	"dropler-new/models"
	"dropler-new/store"
	"log"
)

type Drop struct {
	ID        int64   `json:"id"db:"Id"`
	Name      string  `form:"name"json:"name"`
	Latitude  float32 `form:"lat"json:"lat"`
	Longitude float32 `form:"long"json:"long"`
	Radius    int     `form:"radius"json:"radius"`
	DropGeom  string  `form:"-"json:"-"db:",transient"`
	models.TimeStamp
}

type DropList []Drop

func (d *DropList) List() error {
	_, err := store.Db.Select(d, "SELECT * FROM drops ORDER BY CreatedAt DESC")
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
	err := store.Db.Insert(d)
	if err != nil {
		return err
	}

	return nil
}

func (d *Drop) GetById(id string) error {
	err := store.Db.SelectOne(d, "SELECT * FROM drops WHERE id=$1", id)
	if err != nil {
		return err
	}

	return nil
}
