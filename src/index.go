package src

import (
	// "fmt"
	"gin-gorm/config"
	"gin-gorm/src/database"
	"gin-gorm/src/utils/logger"

	// "time"

	log "github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"flag"
)

func BoostrapApp() {
	// testing

	err := godotenv.Load()

	if err != nil {
		logger.Fatal("Error loading .env file")
	}

	appCfg := config.GetConfig().App
	if appCfg.Mode == "debug" {
		// gin.SetMode(gin.DebugMode)
		// log.SetLevel(log.DebugLevel)
        logger.Log.SetLevel(log.DebugLevel)
		logger.Debug("debug mode")
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	database.ConnectDatabase()
	noDbs := flag.Bool("no-migrate", false, "set to not migrate database.")
	flag.Parse()

	if !*noDbs {

		database.Migrate()
	} else {
		logger.Info("database set to not migrate")

	}
	app := gin.New()

	app.Use(gin.Recovery())

	InitRoute(app)

	if err := app.Run(appCfg.Port); err != nil {
		logger.Fatal("[ERR] fail starting servr :", "err", err)
	}

	logger.Info("[RUN] server run on port", appCfg.Port)
}
