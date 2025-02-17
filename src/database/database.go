package database

import (
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

	// if dbCfg.Driver == "mysql" {
		// dsnMysql := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbCfg.User, dbCfg.Password, dbCfg.Host, dbCfg.Port, dbCfg.Name)
		DB, errConnection = gorm.Open(mysql.Open(dbCfg.DSN), &gorm.Config{})
	// }

	if errConnection != nil {
		logger.Fatal("Cant connect to database")
	}
	// log.Println("connected to database")
	logger.Info("Success connected to database", "driver")

}

func Migrate() {
	err := DB.AutoMigrate(
		&dao.UserEntity{}, 
		&dao.BookEntity{}, 
		&dao.AuthEntity{},
	)

	if err != nil {
		logger.Fatal("Migration Failed", "err", err.Error())
	}
	logger.Info("Database migrated successfully.")
}
