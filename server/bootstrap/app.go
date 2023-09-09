package bootstrap

import "github.com/zag13/tophub/server/dal"

type Application struct {
	Config *Config
	Q      *dal.Query
}

func App() Application {
	app := Application{}

	app.Config = NewConfig()
	app.Q = NewDatabase(app.Config)

	initEnv()

	return app
}
