package main

import (
	"fmt"
	"os"
	"time"

	_ "main.go/docs"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
	viper "github.com/spf13/viper"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Wadzpay go API
// @version         1.0
// @description     This is a sample wadzpay project.
// @termsOfService  http://swagger.io/terms/
// @contact.name   API Support - Blockchain Wadzpay
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:9093
// @BasePath /
// @schemes http

func main() {

	rout := gin.Default()

	rout.Use(Logger()) //use of middleware
	rout.GET("/pagotoken", GetAuthToken)
	rout.GET("/pagoentities", GetListFromPagoEntities)
	rout.GET("/metrics", gin.WrapH(promhttp.Handler()))
	rout.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	log.WithFields(log.Fields{
		"attrib1": "val1",
		"attrib2": 10,
	}).Info("info msg from logger")
	port := viper.GetString("port")
	//prodport := viper.Get("prod.port")

	//servererr := http.ListenAndServe(":"+port, rout)
	//log.Fatal(servererr)

	if err := rout.Run("localhost:" + port); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("current running port :", port)
	}
}

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.InfoLevel)
	// config settings
	viper.AddConfigPath("./configs")
	viper.SetConfigName("config") // Register config file name (no extension)
	viper.SetConfigType("json")   // Look for specific type
	viper.ReadInConfig()

}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// Set example variable
		c.Set("example", "12345")

		// before request

		c.Next()

		// after request
		latency := time.Since(t)
		log.Print(latency, " time latency")

		// access the status we are sending
		status := c.Writer.Status()
		log.Println(status, " as status")
	}
}
