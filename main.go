package main

/* ========================

===========================*/
import (
	"flag"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/pslpune/golang-jumpstart/auth"
	log "github.com/sirupsen/logrus"
)

var (
	FVerbose, FLogF, FSeed bool
	logFile                string
)

func init() {
	/* -------------
	Setting up log configuration for the api
	----------------*/
	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: false,
		ForceColors:   true,
		PadLevelText:  true,
	})
	log.SetReportCaller(false)
	// By default the log output is stdout and the level is info
	log.SetOutput(os.Stdout)     // FLogF will set it main, but dfault is stdout
	log.SetLevel(log.DebugLevel) // default level info debug but FVerbose will set it main
	logFile = os.Getenv("LOGF")
}

func HandlLogin(c *gin.Context) {
	u := auth.NewUser("someone", "some.one@example.com", "timbaktoo", "9845353=3453")
	if u == nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	log.Debug(u.Details())
	yes, _ := u.(auth.Auth).Exists()
	if !yes {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	yes, err := u.(auth.Auth).Login("examplepassword")
	if err != nil || !yes {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
}

func main() {
	flag.Parse() // command line flags are parsed
	log.WithFields(log.Fields{
		"verbose": FVerbose,
		"flog":    FLogF,
		"seed":    FSeed,
	}).Info("Log configuration..")
	if FVerbose {
		log.SetLevel(log.DebugLevel)
	}
	if FLogF {
		lf, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0664)
		if err != nil {
			log.WithFields(log.Fields{
				"err": err,
			}).Error("Failed to connect to log file, kindly check the privileges")
		} else {
			log.Infof("Check log file for entries @ %s", logFile)
			log.SetOutput(lf)
		}
	}
	log.Info("Now starting the telegram scraper microservice")
	gin.SetMode(gin.DebugMode)
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"app":    "Demo application",
			"author": "niranjan_awati@persistent.com",
			"date":   "December 2023",
			"msg":    "If you can see this, its probably running fine",
		})
	})
	r.POST("/users/:id", HandlLogin)
	log.Fatal(r.Run(":8080"))
}
