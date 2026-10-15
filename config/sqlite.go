package config

import (
	"os"

	"github.com/JoaoAnunciacaoDev/Go-Talent-Hub/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	db_path := "./db/main.db"

	// Check if DB already exists
	_, err := os.Stat(db_path)
	if os.IsNotExist(err) {
		logger.Info("database file not founded, creating...")
		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			return nil, err
		}

		file, err := os.Create(db_path)
		if err != nil {
			return nil, err
		}

		file.Close()
	}

	// Create DB and Connect
	db, err := gorm.Open(
		sqlite.Open("./db/main.db"),
		&gorm.Config{},
	)

	if err != nil {
		logger.Errf("sqlite opening error: %v", err)
		return nil, err
	}

	// Migrate the schema
	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errf("sqlite automigration error: %v", err)
		return nil, err
	}

	return db, err
}
