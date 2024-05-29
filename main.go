package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

func add(x, y int) int {
	return x + y
}

func main() {
	r := gin.New()
	r.SetFuncMap(template.FuncMap{
		"add": add,
	})
	r.LoadHTMLGlob("template/*")
	r.GET("/grammar", func(c *gin.Context) {
		c.HTML(http.StatusOK, "grammar", gin.H{
			"hello": "Hello World",
			"map":   map[string]string{"key1": "value1", "key2": "value2"},
			"slice": []string{"slice1", "slice2", "slice3"},
		})
	})
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8888")
}
