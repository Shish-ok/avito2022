package main

import (
	"avito2022/internal/app/api"
	"avito2022/internal/app/config"
	"avito2022/internal/app/storage"
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(
			context.Background(),
			config.NewConfig(),
			gin.Default(),
			storage.NewPostgresStorage,
		),
		fx.Invoke(api.StartHook),
	).Run()
}
