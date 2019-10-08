package infrastructure

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Database struct {
	connection *gorm.DB
}

var singletonDb *Database

func ConnectDatabase() *Database {

	if singletonDb == nil {
		gormDb, err := gorm.Open("sqlite3", "/tmp/gorm.db")
		gormDb.DB().SetMaxIdleConns(10)
		gormDb.DB().SetMaxOpenConns(20)
		gormDb.DB().SetConnMaxLifetime(time.Hour)
		gormDb.LogMode(true)

		if err != nil {
			panic("Failed to open the SQLite database.")
		}
		singletonDb = &Database{connection: gormDb}
	}

	return singletonDb
}

func (db Database) InitializeDatabase(tables ...interface{}) {
	db.connection.AutoMigrate(tables...)
}

func (db Database) GetConnection() *gorm.DB {
	fmt.Println(&singletonDb)
	fmt.Println(&db)
	return db.connection
}

func (db Database) CloseConnection() {
	fmt.Println("Database closed")
	db.connection.Close()
}
