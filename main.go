package main

import (
	"flag"
	"net/http"

	"github.com/uber-go/zap"

	"okr-service/routes"
	"recipe"
	"recipe/config"
)

var logger zap.Logger

func init() {
	recipe.NewLogger(zap.New(
		zap.NewTextEncoder(zap.TextNoTime()),
		zap.DebugLevel,
	))
}

func main() {

	//custom flags...
	flag.Parse()

	// TODO read flags
	conf, err := config.InitConfig("config.yaml")
	if err != nil {
		recipe.Logger.Fatal("configuration error",
			zap.Error(err),
		)
	}

	// add custom routes to default routes.
	recipe.BaseRoutes = recipe.BaseRoutes.AddRoutes(routes.AppRoutes)

	router := recipe.NewRouter(conf.BasePath)
	recipe.Logger.Info("Starting " + conf.ServiceName + " at " + conf.Host + ":" + conf.Port)
	recipe.Logger.Fatal("Error during service", zap.Error(http.ListenAndServe(conf.Host+":"+conf.Port, router)))
}
