package oauth

import (
	"dropler/users"
	"fmt"

	"github.com/RangelReale/osin"
	"github.com/gin-gonic/gin"
)

func Token(c *gin.Context) {
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

	var u users.User

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
