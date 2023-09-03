package config

import (
	"os"

	"github.com/steixeira93/mark/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitialzeSqlite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbPath := "./db/main.db"

	// Check if the DB exists
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Info("sqlite db not found, creating...")
		//Create the database file and directory
		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			return nil, err
		}
		file, err := os.Create(dbPath)
		if err != nil {
			return nil, err
		}
		file.Close()
	}

	//Create DB and connect
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Errorf("sqlite opening error: %v", err)
		return nil, err
	}
	//Migrate thhe Scema
	err = db.AutoMigrate(&schemas.Appointment{})
	if err != nil {
		logger.Errorf("sqlite migration error: %v", err)
		return nil, err
	}
	// Return the DB
	return db, nil
}
