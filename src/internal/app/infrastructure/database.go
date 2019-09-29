package infrastructure

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Database struct {
	connection gorm.db
}

func NewDatabase() *Database {
	database := new(Database)
	return database
}

func (db *Database) Connect() {

	gormDb, err := gorm.Open("sqlite3", "/tmp/gorm.db")

	if err != nil {
		panic("Failed to open the SQLite database.")
	}

	dbConnection.connection = gormDb
	defer dbConnection.connection.Close()
}

func (db *Database) InitializeDatabase(type ... interface) {
  db.connection.AutoMigrate(interface)
}