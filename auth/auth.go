package auth

import (
	"be-go-bookshelf/models"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func HomeHandler(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "/login")
}

func LoginGetHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{
		"content": "",
	})
}

func LoginPostHandler(c *gin.Context) {
	var credential models.Credentials

	err := c.Bind(&credential)
	if err != nil {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{
			"content": "Binding error",
		})
	}

	if credential.Username != os.Getenv("SUPER_USER") ||
		credential.Password != os.Getenv("SUPER_PASS") {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{
			"content": "Username/password is invalid",
		})
	} else {
		claim := jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 5).Unix(),
			Issuer:    "Rickyslash",
			IssuedAt:  time.Now().Unix(),
		}

		sign := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

		secret := os.Getenv("SUPER_SECRET")

		token, err := sign.SignedString([]byte(secret))
		if err != nil {
			c.HTML(http.StatusInternalServerError, "login.html", gin.H{
				"content": "Token signing error",
			})
			c.Abort()
		}

		q := url.Values{}                                              // creates new empty set query parameters
		q.Set("auth", token)                                           // set "auth" query parameter value to "token"
		location := url.URL{Path: "/books", RawQuery: q.Encode()}      // creates url with path `/books` & include the query parameters
		c.Redirect(http.StatusMovedPermanently, location.RequestURI()) // tell server to redirect to the URL, that has been made
	}
}
