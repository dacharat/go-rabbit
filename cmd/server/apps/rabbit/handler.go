package rabbit

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PublishHandler(c *gin.Context) {
	err := Publish()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.Status(http.StatusNoContent)
}
