package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DBConnect() *gorm.DB {
	dsn := "host=localhost user=postgres password=penc3r4han dbname=postgres port=5432 sslmode=disable TimeZone=UTC"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if nil != err {
		log.Printf("ERROR: failed dbConnect open postgres, message: %s", err.Error())
		log.Fatal(err)
	}

	return db
}

// func DBConnect() *gorm.DB {
// 	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%s&loc=%s",
// 		os.Getenv("DATABASE_USER"),
// 		os.Getenv("DATABASE_PASSWORD"),
// 		os.Getenv("DATABASE_HOST"),
// 		os.Getenv("DATABASE_PORT"),
// 		os.Getenv("DATABASE_NAME"),
// 		os.Getenv("DATABASE_CHARSET"),
// 		os.Getenv("DATABASE_PARSE_TIME"),
// 		os.Getenv("DATABASE_LOC"))
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

// 	if nil != err {
// 		log.Printf("ERROR: failed dbConnect open mysql, message: %s", err.Error())
// 		log.Fatal(err)
// 	}

// 	return db
// }
