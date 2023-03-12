package data

import (
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"tophub/app/topdata/cronjob/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewDB)

// Data .
type Data struct {
	db  *gorm.DB
	log *log.Helper
}

// NewData .
func NewData(db *gorm.DB, logger log.Logger) (*Data, func(), error) {
	log := log.NewHelper(log.With(logger, "module", "topdata-cronjob/data"))

	d := &Data{
		db:  db,
		log: log,
	}
	return d, func() {

	}, nil
}

func NewDB(conf *conf.Data, logger log.Logger) *gorm.DB {
	log := log.NewHelper(log.With(logger, "module", "topdata-cronjob/data/gorm"))

	db, err := gorm.Open(mysql.Open(conf.Database.Source), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}

	return db
}
