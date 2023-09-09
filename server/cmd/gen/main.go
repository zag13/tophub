package main

import (
	"fmt"

	"github.com/zag13/tophub/server/bootstrap"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath:       "./dal",
		ModelPkgPath:  "./dal/model",
		FieldNullable: true,
	})

	cfg := bootstrap.NewConfig()
	db, err := gorm.Open(mysql.Open(cfg.MySQLDSN))
	if err != nil {
		panic(fmt.Errorf("connect db fail: %w", err))
	}

	g.UseDB(db)

	// generate all table from database
	g.ApplyBasic(g.GenerateAllTable()...)

	g.Execute()
}
