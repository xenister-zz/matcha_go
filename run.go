package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
)

type User struct {
	Name  string `json:"name"`
	Email Email  `json:"email"`
}

type Email struct {
	Name       string `json:"Owner"`
	Hat        bool   `json:"Has_@"`
	DomaineNbr int    `json:"Domaine_number"`
}

var jresponse string

func main() {
	p := fmt.Println
	//pf := fmt.Printf
	//fmt.Println("Run Api")
	//api.Prter("Prter OK")
	email := Email{Name: "saul", Hat: true, DomaineNbr: 31}

	user := User{Name: "Saul", Email: email}
	//p("Println:")
	//p(user)
	//p("\n")
	//p("Printf:")
	//pf("%+v\n", user)

	byteArray, err := json.Marshal(user)
	if err != nil {
		p(err)
	}
	jresponse = string(byteArray)
	p(jresponse)
	//p("\n")

	r := gin.Default()
	r.GET("/", HomePage)
	r.GET("/query", QueryString)              // Format : "/query?name=tata&age=31"
	r.GET("/path/:name/:age", PathParameters) // Format : "/path/tata/31"
	r.POST("/", PostHomePage)

	r.Run(":81")

}

func PathParameters(c *gin.Context) {
	name := c.Param("name")
	age := c.Param("age")

	c.JSON(200, gin.H{
		"age":  age,
		"name": name,
	})
}

func HomePage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func PostHomePage(c *gin.Context) {
	//body := c.Request.Body
	//value, err := ioutil.ReadAll(body)
	//if err != nil {
	//	fmt.Println("Error Post Home Page")
	//}

	c.JSON(200, gin.H{
		"Message": jresponse,
	})
}

func QueryString(c *gin.Context) {
	name := c.Query("name")
	age := c.Query("age")
	fmt.Println(name, age)

	c.JSON(200, gin.H{
		"age":  age,
		"name": name,
	})

}
