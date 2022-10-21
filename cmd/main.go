package main

import (
	"avito2022/internal/app/config"
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
		),
	).Run()
}
