package main

import (
	b "gin-gorm/src"
	// logrus "github.com/sirupsen/logrus"
)

func main() {
	// logrus.SetFormatter(&logrus.TextFormatter{})
    // logrus.SetLevel(logrus.InfoLevel)
	b.BoostrapApp()
	
}
