package infrastructre

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	dsn := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@" + os.Getenv("DB_HOST") + "/" + os.Getenv("DB_NAME") + "?charset=utf8mb4&parseTime=True&loc=Asia%2FTokyo"
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		panic(err.Error)
	}

	return db
}