package bootstrap

import (
	"github.com/zag13/tophub/server/dal/query"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func NewDatabase(config Config) *query.Query {
	db, err := gorm.Open(mysql.Open(config.MySQLDSN), &gorm.Config{})
	if err != nil {
		log.Fatal("Can't connect to database: ", err)
	}

	return query.Use(db)
}
