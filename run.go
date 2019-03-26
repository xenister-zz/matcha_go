package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xenister/matcha_go/api"
)

func main() {
	fmt.Println("Run Api")
	api.Prter("Prter OK")

	r := gin.Default()
	r.GET("/", HomePage)
	r.GET("/query", QueryString) // Format : "/query?name=tata&age=31"
	r.POST("/", PostHomePage)

	r.Run(":81")

}
func HomePage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func PostHomePage(c *gin.Context) {
	c.JSON(200, gin.H{
		"Message": "Post Home Page",
	})
}

func QueryString(c *gin.Context) {
	name := c.Query("name")
	age := c.Query("age")
	fmt.Println(name, age)

	c.JSON(200, gin.H{
		"pge":  age,
		"name": name,
	})
}
