package main

import (
	"fmt"

	"github.com/zag13/tophub/server/bootstrap"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func init() {
	cfg := bootstrap.NewConfig()
	db, err = gorm.Open(mysql.Open(cfg.MySQLDSN))
	if err != nil {
		panic(fmt.Errorf("connect db fail: %w", err))
	}

	prepare(db) // prepare table for generate
}

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath:       "./dal",
		ModelPkgPath:  "./dal/model",
		FieldNullable: true,
	})

	g.UseDB(db)

	// g.ApplyBasic(g.GenerateAllTable()...)
	g.ApplyInterface(func(Querier) {}, g.GenerateAllTable()...)

	g.Execute()
}
