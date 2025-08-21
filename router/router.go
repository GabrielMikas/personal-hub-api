package router

import "github.com/gin-gonic/gin"

func Run() {
	g := gin.Default()
	serveRoutes(g)
	g.Run(":3000")
}
