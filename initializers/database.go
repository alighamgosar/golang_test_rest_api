package initializers

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConntectToDB() {
	var err error
	// postgres://fekhttca:67uBlhK2ZbiZL6JoUsC4OjKMqO4q1Y8T@silly.db.elephantsql.com/fekhttca
	dsn := os.Getenv("DB_URL")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// _ = DB

	if err != nil {
		log.Fatal("Failed to connect to database")
	}
}
