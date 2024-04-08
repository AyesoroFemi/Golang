package initializers

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


var DB *gorm.DB
func ConnectToDb() {
				// username password-------------
	// postgres://dsiiuooe:o_qU8Abs4U1Bge-WZv47fFbavo0uo33i@surus.db.elephantsql.com/dsiiuooe
	var err error
	dsn := os.Getenv("DB_URL")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	// if err != nil {
	// 	panic("Failed to connect to db")
	// }
	if err != nil {
		log.Fatal("Failed to connect to database")
	}

}

