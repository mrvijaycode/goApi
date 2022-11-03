package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
	viper "github.com/spf13/viper"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	//_ "wadzpay.com/poc/app1/docs"
)

// @title           Swagger Example API
// @version         2.0
// @description     This is a sample helloworld project.
// @termsOfService  http://swagger.io/terms/
// @contact.name   API Support - Blockchain Wadzpay
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	/* CORS
	   "github.com/gin-contrib/cors"
	    router.Use(cors.Default())
	*/
	rout := gin.Default()

	rout.Use(Logger()) //use of middleware

	//rout.GET("/helloworld", Logger(), Greet)
	//rout.GET("/helloworld", Greet)
	rout.GET("/pagotoken", GetAuthToken)
	rout.GET("/pagoentities", GetListFromPagoEntities)
	rout.GET("/metrics", gin.WrapH(promhttp.Handler()))
	rout.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	log.WithFields(log.Fields{
		"attrib1": "val1",
		"attrib2": 10,
	}).Info("info msg from looger")
	port := viper.GetString("port")
	//prodport := viper.Get("prod.port")

	fmt.Println("current running port :", port)
	//fmt.Println("current running port for prod :", prodport)
	// rout.Run("localhost:8080") // assumes localhost:8080
	servererr := http.ListenAndServe(":"+port, rout)
	log.Fatal(servererr)
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
	/*
			viper.WatchConfig()
		viper.OnConfigChange(func(e fsnotify.Event) {
		    fmt.Println("Config file changed:", e.Name)
		})
			//
			// The WatchConfig() method will look for changes in the config file,
			// while the onConfigChange is an optional method that will run each time the file changes
	*/

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

/*
 rout.Use(FindUserAgent())
func FindUserAgent() gin.HandlerFunc {
    return func(c *gin.Context) {
        log.Println(c.GetHeader("User-Agent"))
        // Before calling handler
        c.Next()
        // After calling handler
    }
}
*/
