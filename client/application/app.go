package application

import (
	connect "github.com/ankitanwar/assignment/client/ServerConnect"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	urlMapping()
	connect.ConnectServer()
	router.Run("0.0.0.0:8070")
}
