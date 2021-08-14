package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func EndpointStarter() {
	var r = gin.Default()
	r.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "this is get method!")
	})

	r.Run(":8080")
}
