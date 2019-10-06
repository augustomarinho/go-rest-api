package infrastructure

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Database struct {
	connection *gorm.DB
}

func NewDatabase() *Database {

	database := new(Database)
	database.connect()
	return database
}

func (db *Database) connect() {

	gormDb, err := gorm.Open("sqlite3", "/tmp/gorm.db")
	gormDb.DB().SetMaxIdleConns(10)
	gormDb.DB().SetMaxOpenConns(20)
	gormDb.DB().SetConnMaxLifetime(time.Hour)
	gormDb.LogMode(true)

	if err != nil {
		panic("Failed to open the SQLite database.")
	}

	db.connection = gormDb

}

func (db *Database) InitializeDatabase(tables ...interface{}) {
	db.connection.AutoMigrate(tables...)
}

func (db *Database) GetConnection() *gorm.DB {
	return db.connection
}

func (db *Database) CloseConnection() {
	db.connection.Close()
}
