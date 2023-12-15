package main

import (
	logger "goGinBackend/Logger"
	middleware "goGinBackend/Middleware"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/mattn/go-colorable"
	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

func main() {
	logrus.SetLevel(logrus.TraceLevel)

	logrus.SetReportCaller(true)

	logrus.SetFormatter(&logrus.JSONFormatter{
		DisableTimestamp: true,
		PrettyPrint:      true,
	})

	f, _ := os.Create("logrus.log")

	multi := io.MultiWriter(f, os.Stdout)

	logrus.SetOutput(multi)
	logrus.Traceln("trace")
	logrus.Debugln("debug")
	logrus.Infoln("info")
	logrus.Warnln("warn")
	logrus.Errorln("error")
	// logrus.Panicln("panic")

	router := gin.New()

	logrus.Println("Hi I am logrus")

	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Printf("endpoint formatted information is %v %v %v %v/n", httpMethod, absolutePath, handlerName, nuHandlers)
	}

	gin.ForceConsoleColor()
	gin.DefaultWriter = colorable.NewColorableStdout()
	// router.Use(middleware.Authenticate) // pass to the whole application
	f1, _ := os.Create("ginLogging.Log")
	gin.DefaultWriter = io.MultiWriter(f1, os.Stdout)

	router.Use(gin.LoggerWithFormatter(logger.FormatLogs))
	router.GET("/getData1", middleware.Authenticate, middleware.AddHeader, getData1)

	auth := gin.BasicAuth(gin.Accounts{
		"user": "pass",
	})

	admin := router.Group("/admin", auth)
	{
		admin.GET("/getData", getData)
	}

	client := router.Group("/client", middleware.Authenticate)
	{
		client.GET("/getQueryString", getQueryString)
		client.GET("/getData2", getData2)
		client.GET("/getData3", getData3)
	}

	// router.GET("/getData", getData)
	// router.GET("/getQueryString", getQueryString)
	router.GET("/getUrlData/:name/:age", getUrlData)
	router.POST("/getDataPost", getDataPost)
	router.Run(":5000")
	// server := &http.Server{
	// 	Addr:         ":5000",
	// 	Handler:      router,
	// 	ReadTimeout:  10 * time.Second,
	// 	WriteTimeout: 10 * time.Second,
	// }
	// server.ListenAndServe()
}

func getData(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": "Hi I am GIN Framework",
	})
}

func getData1(c *gin.Context) {

	logrus.WithField("Info", "CreateFile").Info("starting file creation")
	logrus.WithField("Debug", "CreateFile").Debug("Starting File Creation")
 
	f, err:= os.Create("dhdhdj/dddjjd/sdkd/logrus.log")

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"Method" : "CreateFile",
			"error" : true,
		}).Error(err.Error())
	}

    logrus.WithField("Debug", "CreateFile").Debug("Starting File Creation")

	multi := io.MultiWriter(f, os.Stdout)
	logrus.SetOutput(multi)
	
	c.JSON(200, gin.H{
		"data": "Hi I 1 am GIN Framework",
	})
}

func getData2(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": "Hi I 2 am GIN Framework",
	})
}

func getData3(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": "Hi I 3 am GIN Framework",
	})
}

// http://localhost:5000/getQueryString?name=James&age=15
func getQueryString(c *gin.Context) {
	name := c.Query("name")
	age := c.Query("age")
	c.JSON(200, gin.H{
		"data": "Hi I am GIN Framework",
		"name": name,
		"age":  age,
	})
}

// http://localhost:5000/getUrlData/mark/14
func getUrlData(c *gin.Context) {
	name := c.Param("name")
	age := c.Param("age")
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

//Trace
//Debug
//Info
//Warn
//Error
//Fatal
//Panic
