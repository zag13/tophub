package main

import (
	"fmt"

	"github.com/zag13/tophub/cli/bootstrap"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func init() {
	bootstrap.InitConfig()
	db, err = gorm.Open(mysql.Open(bootstrap.C.MySQLDSN))
	if err != nil {
		panic(fmt.Errorf("connect db fail: %w", err))
	}

	prepare(db)
}

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath:       "./dal",
		ModelPkgPath:  "./dal/model",
		Mode:          gen.WithDefaultQuery,
		FieldNullable: true,
	})

	g.UseDB(db)

	g.ApplyBasic(g.GenerateAllTable()...)

	g.Execute()
}

func prepare(db *gorm.DB) {
	db.Exec(bootstrap.TableSQL)
}
