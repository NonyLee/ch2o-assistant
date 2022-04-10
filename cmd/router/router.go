package router

import (
	"env/cmd/router/handler"

	"github.com/gin-gonic/gin"
)

func Run() {
	r := gin.Default()
	handler.BuildCh2o(r)

	r.Run(":9001")
}
