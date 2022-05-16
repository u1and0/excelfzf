package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/u1and0/excelfzf/api"
)

func main() {
	r := gin.Default()
	r.Static("/static", "./static")
	r.LoadHTMLGlob("template/*.tmpl")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{}) // エントリポイント
	})
	r.GET("/map", api.FetchMap)   // JSONを配信するAPI
	r.GET("/list", api.FetchList) // JSONを配信するAPI
	r.Run()
}
