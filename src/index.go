package src 

import (
	// "fmt"
	"gin-gorm/config"
	"gin-gorm/src/database"
	"gin-gorm/src/utils/logger"

	// "time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
    log "github.com/charmbracelet/log"
)

func BoostrapApp() {

    err := godotenv.Load()
    
    if err != nil {
        logger.Fatal("Error loading .env file")
    }

    appCfg := config.GetConfig().App
    if appCfg.Mode == "debug" {
        log.SetLevel(log.DebugLevel)
        gin.SetMode(gin.DebugMode)
    } else {
        gin.SetMode(gin.ReleaseMode)
    }

	database.ConnectDatabase()
    database.Migrate()
	app := gin.New()

    app.Use(gin.Recovery())

	InitRoute(app)

	

    if err := app.Run(appCfg.Port); err != nil {
        logger.Fatal("[ERR] fail starting servr :", "err", err)
    } 

    logger.Info("[RUN] server run on port", appCfg.Port)
} 