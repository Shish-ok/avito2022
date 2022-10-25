package main

import (
	"avito2022/internal/app/api"
	"avito2022/internal/app/config"
	"avito2022/internal/app/service/balance"
	"avito2022/internal/app/service/balance_holder"
	"avito2022/internal/app/service/user_transaction"
	"avito2022/internal/app/storage"
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func balanceStorage(storage *storage.PostgresStorage) balance.Storage {
	return storage
}

func userTransactionStorage(storage *storage.PostgresStorage) user_transaction.Storage {
	return storage
}

func balanceHolderStorage(storage *storage.PostgresStorage) balance_holder.Storage {
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
			userTransactionStorage,
			user_transaction.NewService,
			balanceHolderStorage,
			balance_holder.NewService,
		),
		fx.Invoke(api.StartHook),
	).Run()
}
