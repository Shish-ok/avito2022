package api

import (
	"avito2022/docs"
	"avito2022/internal/app/config"
	"avito2022/internal/app/service/balance"
	"avito2022/internal/app/service/user_transaction"
	"context"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/fx"
)

const (
	Title    = "Balance handler API"
	BasePath = "/api/v1/"
)

type Api struct {
	router          *gin.Engine
	balance         *balance.Service
	userTransaction *user_transaction.Service
}

func (api *Api) Run() {
	cfg := config.Load()
	docs.SwaggerInfo.BasePath = BasePath
	docs.SwaggerInfo.Title = Title

	api.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.DefaultModelsExpandDepth(-1)))

	api.router.Use(CORSMiddleware())
	api.router.Run(cfg.Api.GetAddr())
}

func NewApi(
	router *gin.Engine,
	balance *balance.Service,
	userTransaction *user_transaction.Service,
) *Api {
	svc := &Api{
		router:          router,
		balance:         balance,
		userTransaction: userTransaction,
	}
	svc.registerRoutes()
	return svc
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, accept, origin, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func StartHook(lifecycle fx.Lifecycle, api *Api) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				go api.Run()
				return nil
			},
		})
}

func (api *Api) registerRoutes() {
	base := api.router.Group(BasePath)

	balance := base.Group("/balance")
	balance.POST("/up_balance/", api.UpBalance)
	balance.GET("/get_balance/", api.GetBalance)
}
