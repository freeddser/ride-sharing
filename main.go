package main

import (
	"flag"
	"fmt"
	"github.com/freeddser/rs-common/config"
	"github.com/freeddser/rs-common/logging"
	"github.com/freeddser/ride-sharing/repository"
	R "github.com/freeddser/ride-sharing/router"
	"github.com/freeddser/ride-sharing/services"
	"github.com/freeddser/rs-common/util"
	"github.com/jasonlvhit/gocron"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"strconv"
	"time"
)

var log = logging.MustGetLogger()

func main() {
	configFile := flag.String("c", "", "Configuration File")
	flag.Parse()

	if *configFile == "" {
		fmt.Println("\n\nUse -h to get more information on command line options\n")
		fmt.Println("You must specify a configuration file")
		os.Exit(1)
	}

	err := config.Initialize(*configFile)
	if err != nil {
		fmt.Printf("Error reading configuration: %s\n", err.Error())
		os.Exit(1)
	}

	InitRedisFactory()

	util.InitTimeZoneLocation()
	start := util.GetTimeNow()

	setupLogging()

	// Must be initialized before backends
	err = util.InitializeMaps(config.MustGetString("server.maps"))
	if err != nil {
		log.Fatal("error reading map_data.json: ", err.Error())
		return
	}

	//mysql
	if config.MustGetString("switch.mysql") == "on" {
		err = repository.InitMysqlDbFactory()
		if err != nil {
			log.Fatal("Cannot connect to Mysqldatabase: ", err.Error())
			return
		}
	}

	err = services.InitializeService(config.MustGetString("server.mode"))
	if err != nil {
		log.Fatal("Fail to initialize service: ", err.Error())
		return
	}
	//http server
	log.Info("Server Started in ", time.Since(start))
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(config.MustGetInt("Server.port")), R.NewRouter()))

	//cronjob
	s := gocron.NewScheduler()
	<-s.Start()

}

func setupLogging() {
	logrus.SetLevel(logrus.DebugLevel)
	if config.MustGetString("server.mode") == "production" {
		log.Info("here")
		logPath := config.MustGetString("server.log_path")

		file, err := os.OpenFile(logPath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0777)
		if err != nil {
			log.Fatal("Cannot log to file", err.Error())
		}

		logrus.SetFormatter(&logrus.JSONFormatter{})
		logrus.SetOutput(file)
	}
}

func InitRedisFactory() {
	util.InitRedis(config.MustGetString("cache.redis_main"))
}
