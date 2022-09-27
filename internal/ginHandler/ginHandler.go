package ginHandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
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

func Broadcast(c *gin.Context) {
	body := map[string]string{}
	if err := c.ShouldBindBodyWith(&body, binding.JSON); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	lineHandler.Broadcast(body["text"])
	c.JSON(200, "ok")
}

func GetMessagesByUserID(c *gin.Context) {
	userID := c.Param("userID")
	messages := mongoHandler.FindMessagesByUserID(userID)
	c.JSON(200, messages)
}

func GetMessages(c *gin.Context) {
	messages := mongoHandler.FindMessages()
	c.JSON(200, messages)
}