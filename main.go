package main

import (
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/getData", getData)
	router.GET("/getQueryString", getQueryString)
	router.POST("/getDataPost", getDataPost)
	router.Run(":5000")
}

func getData(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": "Hi I am GIN Framework",
	})
}

//http://localhost:5000/getQueryString?name=James&age=15
func getQueryString(c *gin.Context) {
	name := c.Query("name")
	age := c.Query("age")
	c.JSON(200, gin.H{
		"data": "Hi I am GIN Framework",
		"name": name,
		"age":  age,
	})
}
func getDataPost(c *gin.Context) {
	body := c.Request.Body
	value, _ := ioutil.ReadAll(body)
	c.JSON(200, gin.H{
		"data":     "Hi I am GIN post  Framework",
		"bodyData": string(value),
	})
}
