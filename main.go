package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"os"
)

func main() {
	r := gin.New()
	r.LoadHTMLGlob("template/*")
	r.GET("/grammar", func(c *gin.Context) {
		c.HTML(http.StatusOK, "grammar", gin.H{
			"hello": "Hello World",
		})
	})
	parse, err := template.New("").Parse("")
	if err != nil {
		return
	}
	parse.Execute(os.Stdout, "")
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8888")
}
