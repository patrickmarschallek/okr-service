package main

import (
	"fmt"
	"net/http"

	"github.com/uber-go/zap"

	"okr-service/routes"
	"recipe"
	"recipe/config"
)

func init() {
	recipe.Logger = zap.New(zap.NewTextEncoder(zap.TextNoTime()),
		zap.DebugLevel,
	)
}

func main() {
	// defer func() {
	// 	if r := recover(); r != nil {
	// 		recipe.Logger.Fatal("Recoverd error.",
	// 			zap.Object("recover", r),
	// 		)
	// 	}
	// }()
	defer func() {
		if r := recover(); r != nil {
			err, ok := r.(error)
			if !ok {
				err = fmt.Errorf("pkg: %+v", r)
			}
			fmt.Printf("ERROR: %+v", err)
		}
	}()

	conf, err := config.ReadConfig("config.yaml")
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
