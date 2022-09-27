package ginHandler

import (
	"github.com/gin-gonic/gin"
	"github.com/kzeratal/cinnox-homework/internal/lineHandler"
	"github.com/kzeratal/cinnox-homework/internal/mongoHandler"
)

func Revceive(c *gin.Context) {
	messages := lineHandler.GetMessages(c.Request)
	if len(messages) > 0 {
		mongoHandler.InsertMany(messages)
	}
	c.JSON(200, "ok")
}