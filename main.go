package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("resources/*.templ.html")
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "login.templ.html", gin.H{})
	})
	r.Run()
}
