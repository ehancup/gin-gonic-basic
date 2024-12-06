package database

import (
	"fmt"
	// "gin-gorm/app/models"
	"gin-gorm/config"
	"gin-gorm/src/database/dao"
	// "gin-gorm/src/app/users"
	"gin-gorm/src/utils/logger"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {

	var errConnection error
	dbCfg := config.GetConfig().DB

	if dbCfg.Driver == "mysql" {
		dsnMysql := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbCfg.User, dbCfg.Password, dbCfg.Host, dbCfg.Port, dbCfg.Name)
		DB, errConnection = gorm.Open(mysql.Open(dsnMysql), &gorm.Config{})
	}

	if errConnection != nil {
		logger.Fatal("Cant connect to database")
	}
	// log.Println("connected to database")
	logger.Info("connected to database", "driver",dbCfg.Driver )

}

func Migrate() {
	err := DB.AutoMigrate(&dao.UserEntity{})

	if err != nil {
		logger.Fatal("Migration Failed", "err", err.Error())
	}
	logger.Info("Database migrated successfully.")
}