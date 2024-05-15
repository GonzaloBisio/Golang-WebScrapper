package message_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/pusher/pusher-http-go/v5"
	"net/http"
)

func SendMessage(c *gin.Context) {
	pusherClient := pusher.Client{
		AppID:   "1803282",
		Key:     "d43aef0a118705880d2e",
		Secret:  "31455f57fd148b50255c",
		Cluster: "sa1",
		Secure:  true,
	}

	var data map[string]string
	if err := c.ShouldBind(&data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	pusherClient.Trigger("chat", "message", data)
	c.JSON(http.StatusOK, gin.H{"message": "Message sent successfully"})
}
