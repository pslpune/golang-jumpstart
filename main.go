package main

/* ========================
This is a demo project and the only purpose of this code is walk the candidates thru the code structure
While the code does run and starts an HTTP web api server it has no utility whatsoever.
===========================*/
import (
	"flag"
	"net/http"
	"os"
	"time"

	
	"github.com/gin-gonic/gin"
	"github.com/pslpune/golang-jumpstart/auth"
	log "github.com/sirupsen/logrus"
)

const (
	DELAY = 3 * time.Second
)

var (
	Temp =90.00
	FVerbose, FLogF, FSeed bool
	logFile                string
	allUsers               []auth.User = []auth.User{
		&auth.AnyUser{Name: "Querida Fortnam", Email: "qfortnam0@desdev.cn", Loc: "Philippines"},
		&auth.AnyUser{Name: "Vanessa Fay", Email: "vfay1@ning.com", Loc: "Venezuela"},
		&auth.AnyUser{Name: "Christoforo Birdfield", Email: "cbirdfield2@mapy.cz", Loc: "China"},
		&auth.AnyUser{Name: "Beatrix Dottridge", Email: "bdottridge3@alexa.com", Loc: "China"},
		&auth.AnyUser{Name: "Holly Rubert", Email: "hruberti4@cdbaby.com", Loc: "Indonesia"},
		&auth.AnyUser{Name: "Lexy Fendt", Email: "lfendt5@edublogs.org", Loc: "Philippines"},
	} // in memory database of the registerd users
)

func init() {
	/* -------------
	Setting up log configuration for the api
	or reading values from the environment
	global database connections maybe
	global cache connections
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
	log.WithFields(log.Fields{
		"count": len(allUsers),
	}).Debug("database loaded")
	
	
}

func HandlGetUser(c *gin.Context) {
	for idx, usr := range allUsers {
		if email, _ := c.Params.Get("email"); email == usr.(*auth.AnyUser).Email {
			log.WithFields(log.Fields{
				"index": idx,
			}).Debug("user found")
			c.AbortWithStatus(http.StatusOK)
		}
	}
	<-time.After(DELAY) // deliberate delay in sequential process.
	c.AbortWithStatus(http.StatusNotFound)
}

func HandlLogin(c *gin.Context) {
	u := auth.NewUser("someone", "some.one@example.com", "timbaktoo", "98453533453")
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
	claim, _ := c.Params.Get("password")
	yes, err := u.(auth.Auth).Login(claim)
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	if !yes {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	c.AbortWithStatus(http.StatusOK)
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
	r.GET("/users/:email", HandlGetUser)
	log.Fatal(r.Run(":8080"))
}
