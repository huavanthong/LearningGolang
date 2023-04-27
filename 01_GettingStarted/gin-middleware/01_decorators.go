//Golang doesn't have python-Django like decorators but here is
//a small example of what you can do

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/appengine/user"
)

type User struct {
	isAdmin bool
}

func Handler(h gin.HandlerFunc, decors ...func(gin.HandlerFunc) gin.HandlerFunc) gin.HandlerFunc {
	for i := range decors {
		d := decors[len(decors)-1-i] // iterate in reverse
		h = d(h)
	}
	return h
}

func isSuperUser(handlerFunc gin.HandlerFunc) gin.HandlerFunc {
	return func(context *gin.Context) {

		// assuming you already have user through middleware
		if user.IsAdmin {
			handlerFunc(context)
		} else {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "you are not authorized"})
		}
	}
}

func main() {
	router := gin.Default()
	router.GET("/welcome", Handler(welcome, isSuperUser))
}

func welcome(context *gin.Context) {
	context.JSON(200, gin.H{"status": "ok"})
}
