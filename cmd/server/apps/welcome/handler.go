package welcome

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func WelcomeHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"text": "Hello",
	})
}
