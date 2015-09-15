package main

import (
	"droppio/utils/time"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"golang.org/x/crypto/bcrypt"
)

// User model struct
type User struct {
	ID             int64  `json:"id"db:"Id"`
	Name           string `form:"name"json:"name"`
	Email          string `form:"email"json:"email"`
	HashedPassword string `json:"hashed_password"`
	models.TimeStamp
}

type UserList []User

func (u *UserList) List() error {
	_, err := Db.Select(u, "SELECT * FROM USERS ORDER BY CreatedAt DESC")
	if err != nil {
		log.Fatal(err)
	}
	return err
}

// Insert Method to create a new user from the models User struct
func (u *User) Insert(password string) error {

	// Lowercase email
	u.Email = strings.ToLower(u.Email)

	// run the SetPassword method on the user model
	// if a password is provided
	if password != "" {
		err := u.SetPassword(password)
		if err != nil {
			return err
		}
	}

	// run the UpdateTime ethod on the user model
	u.UpdateTime()

	// run the DB insert function
	err := Db.Insert(u)
	if err != nil {
		return err
	}

	return nil
}

func (u *User) GetById(id int64) error {
	err := Db.SelectOne(u, "SELECT * FROM users WHERE id=$1", id)
	if err != nil {
		return err
	}

	return nil
}

// SetPassword Method on User model for setting the hashed password
func (u *User) SetPassword(password string) error {

	b, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.HashedPassword = string(b)

	return nil
}

// CheckPassword Method to check if password matches the stored hash
func (u *User) CheckPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.HashedPassword), []byte(password))
}

// FindByEmail Method for returning a single user row from
// a provided email param.
func (u *User) FindByEmail(email string) error {

	err := Db.SelectOne(u, "select * from users where email=$1", email)

	return err
}

// List Method for returning all users as array of objects
func ListUser(c *gin.Context) {
	fmt.Println("In List Controller")
	u := UserList{}

	err := u.List()
	if err != nil {
		c.JSON(500, gin.H{"error": "Problem fetching users. No database connection?"})
	}

	c.JSON(200, u)
}

func CreateUser(c *gin.Context) {
	u := User{}

	c.BindWith(&u, binding.Form)

	// Get password from form value directly and
	// not through the BindWith method.
	password := c.Request.FormValue("password")

	err := u.Insert(password)
	if err != nil {
		c.JSON(500, gin.H{"error": "Problem fetching users. No database connection?"})
		return
	}

	c.JSON(200, u)
}

func GetUser(c *gin.Context) {
	u := User{}
	id := c.Params.ByName("id")

	intid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(500, err)
		return
	}

	err = u.GetById(intid)
	if err != nil {
		c.JSON(500, gin.H{"error": "Problem fetching users. No database connection?"})
		return
	}

	c.JSON(200, u)
}
