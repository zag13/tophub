package bootstrap

import (
	"log"
	"sync"

	"github.com/zag13/tophub/cli/dal"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var onceDb sync.Once

func InitDatabase() {
	onceDb.Do(func() {
		db, err := gorm.Open(mysql.Open(C.MySQLDSN), &gorm.Config{})
		if err != nil {
			log.Fatal("Can't connect to database: ", err)
		}

		dal.SetDefault(db)
	})
}
