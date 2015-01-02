package oauth

import (
	"dropler/users"
	"fmt"
	"github.com/RangelReale/osin"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
)

func Authorize(c *gin.Context) {
	var (
		oauth = OauthServer
		resp  = oauth.NewResponse()
	)

	// Defer the oauth response close function
	defer resp.Close()

	if ar := oauth.HandleAuthorizeRequest(resp, c.Request); ar != nil {
		if !handleLoginPage(ar, c.Writer, c.Request) {
			return
		}

		ar.UserData = struct{ Login string }{Login: "test"}
		ar.Authorized = true
		oauth.FinishAuthorizeRequest(resp, c.Request, ar)
	}

	if resp.IsError && resp.InternalError != nil {
		fmt.Printf("ERROR: %s\n", resp.InternalError)
	}

	// User osin json output for now. It sets headers and status code
	osin.OutputJSON(resp, c.Writer, c.Request)

}

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

// HandleLoginPage temp func
func handleLoginPage(ar *osin.AuthorizeRequest, w http.ResponseWriter, r *http.Request) bool {
	r.ParseForm()
	if r.Method == "POST" && r.Form.Get("login") == "test" && r.Form.Get("password") == "test" {
		return true
	}

	w.Write([]byte("<html><body>"))

	w.Write([]byte(fmt.Sprintf("LOGIN %s (use test/test)<br/>", ar.Client.GetId())))
	w.Write([]byte(fmt.Sprintf("<form action=\"/authorize?response_type=%s&client_id=%s&state=%s&redirect_uri=%s\" method=\"POST\">",
		ar.Type, ar.Client.GetId(), ar.State, url.QueryEscape(ar.RedirectUri))))

	w.Write([]byte("Login: <input type=\"text\" name=\"login\" /><br/>"))
	w.Write([]byte("Password: <input type=\"password\" name=\"password\" /><br/>"))
	w.Write([]byte("<input type=\"submit\"/>"))

	w.Write([]byte("</form>"))

	w.Write([]byte("</body></html>"))

	return false
}
