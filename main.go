// main.go
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Student struct {
	Name string
	Age  int8
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	stu1 := &Student{Name: "Tom", Age: 19}
	stu2 := &Student{Name: "Sam", Age: 30}

	r.GET("/user", func(c *gin.Context) {
		name := c.Query("name")
		role := c.DefaultQuery("role", "teacher")
		c.String(http.StatusOK, "%s is a %s", name, role)
	})
	r.POST("/form", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")

		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"password": password,
		})
	})

	r.GET("/arr", func(c *gin.Context) {
		c.HTML(http.StatusOK, "arr.tmpl", gin.H{
			"title":  "Gin",
			"stuArr": [2]*Student{stu1, stu2},
		})
	})
	r.Run(":9999") // listen and serve on 0.0.0.0:8080
}
