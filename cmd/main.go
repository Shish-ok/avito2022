package main

import (
	"avito2022/internal/app/api"
	"avito2022/internal/app/config"
	"avito2022/internal/app/service/balance"
	"avito2022/internal/app/storage"
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func balanceStorage(storage *storage.PostgresStorage) balance.Storage {
	return storage
}

// main godoc
func main() {
	fx.New(
		fx.Provide(
			context.Background,
			config.NewConfig,
			api.NewApi,
			gin.Default,
			storage.NewPostgresStorage,
			balanceStorage,
			balance.NewService,
		),
		fx.Invoke(api.StartHook),
	).Run()
}
