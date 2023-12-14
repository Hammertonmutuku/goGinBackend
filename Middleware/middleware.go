package middleware

import "github.com/gin-gonic/gin"

//way one to write middleware
// func Authenticate() gin.HandlerFunc {

// 	// write custom  logic to be applied beofre my middleware is executed

// 	return func(c *gin.Context) {
// 		if !(c.Request.Header.Get("Token") == "auth") {
// 			c.AbortWithStatusJSON(500, gin.H{
// 				"Message": "Token Not Present",
// 			})
// 			return
// 		}
// 		c.Next()
// 	}
// }

//way two to write middleware
func Authenticate(c *gin.Context) {
	if !(c.Request.Header.Get("Token")== "auth") {
		c.AbortWithStatusJSON(500, gin.H{
			"Message" : "Token Not Present",
		})
		return
	}
	c.Next()
}

func AddHeader(c *gin.Context) {
	c.Writer.Header().Set("Key", "Value")
	c.Next()
}
