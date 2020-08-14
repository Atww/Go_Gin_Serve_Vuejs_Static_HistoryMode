package main

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	r.LoadHTMLGlob("./static/dist/*.html") // load the built dist path
	r.LoadHTMLFiles("./static/*")          //  load the static path

	r.Use(static.Serve("/tracking/", static.LocalFile("./static/dist", true)))
	r.NoRoute(func(c *gin.Context) {
		c.File("./static/dist/index.html")
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
