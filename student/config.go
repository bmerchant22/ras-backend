package student

import (
	"github.com/sirupsen/logrus"

	"github.com/spf13/viper"
	_ "github.com/spo-iitk/ras-backend/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func openConnection() {
	host := viper.GetString("DATABASE.HOST")
	port := viper.GetString("DATABASE.PORT")
	password := viper.GetString("DATABASE.PASSWORD")

	dbName := viper.GetString("DBNAME.STUDENT")
	user := dbName + viper.GetString("DATABASE.USER")

	dsn := "host=" + host + " user=" + user + " password=" + password
	dsn += " dbname=" + dbName + " port=" + port + " sslmode=disable TimeZone=Asia/Kolkata"

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		logrus.Fatal("Failed to connect to student database: ", err)
		panic(err)
	}

	db = database

	err = db.AutoMigrate(&Student{}, &StudentDocument{})
	if err != nil {
		logrus.Fatal("Failed to migrate student database: ", err)
		panic(err)
	}

	logrus.Info("Connected to student database")
}

func init() {
	openConnection()
}
