package bootstrap

import "github.com/zag13/tophub/server/dal/query"

type Application struct {
	Config Config
	Q      *query.Query
}

func App() Application {
	app := &Application{}
	app.Config = NewConfig()
	app.Q = NewDatabase(app.Config)
	return *app
}
