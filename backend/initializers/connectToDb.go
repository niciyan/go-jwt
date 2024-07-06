package initializers

import (
	"gorm.io/driver/sqlite" // Sqlite driver based on CGO
	// "github.com/glebarez/sqlite" // Pure go SQLite driver, checkout https://github.com/glebarez/sqlite for details
	"gorm.io/gorm"
)

  var DB *gorm.DB
  
  func ConnectToDb() {

	  // github.com/mattn/go-sqlite3
	  db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	  if err != nil {
		panic(err)
	  }
	  DB = db
  }