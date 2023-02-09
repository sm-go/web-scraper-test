package ds

import (
	"fmt"
	"log"
	"os"

	"github.com/smith-golang/Sites/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

type DataSource struct {
	DB *gorm.DB
}

func NewDataSource() (*DataSource, error) {
	db, err := LoadDB()
	if err != nil {
		return nil, err
	}
	DB = db
	return &DataSource{
		DB: db,
	}, nil
}

func LoadDB() (*gorm.DB, error) {
	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")
	user := os.Getenv("MYSQL_USER")
	pass := os.Getenv("MYSQL_PASS")
	name := os.Getenv("MYSQL_NAME")

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, pass, host, port, name,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(
		&model.Post{},
	)
	if err != nil {
		log.Println(err)
	}

	return db, nil
}
