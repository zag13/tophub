package main

import (
	"github.com/zag13/tophub/server/bootstrap"
	"gorm.io/gen"
	"gorm.io/gen/examples/dal"
)

func init() {
	app := bootstrap.App()
	dal.DB = dal.ConnectDB(app.Config.MySQLDSN).Debug()
}

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath:       "./dal/query",
		Mode:          gen.WithDefaultQuery | gen.WithoutContext | gen.WithQueryInterface,
		FieldNullable: true,
	})

	g.UseDB(dal.DB)

	// generate all table from database
	g.ApplyBasic(g.GenerateAllTable()...)

	g.Execute()
}
