package main

import (
	b "gin-gorm/src"
	// logrus "github.com/sirupsen/logrus"
)

//	@title			Learn Golang gin-gonic REST API
//	@version		1.0
//	@description	This is an API for learning.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	Ehancup
//	@contact.url	http://www.swagger.io/support
//	@contact.email	rhanysuf24@gmail.com

//	@host		localhost:3010
//	@BasePath	/v1

//	@securityDefinitions.apikey	BearerAuth
//	@in							header
//	@name						Authorization
func main() {
	// logrus.SetFormatter(&logrus.TextFormatter{})
    // logrus.SetLevel(logrus.InfoLevel)
	b.BoostrapApp()
	
}
