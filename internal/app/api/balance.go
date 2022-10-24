package api

import (
	"avito2022/internal/app/models"
	balance_service "avito2022/internal/app/service/balance"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserID struct {
	ID uint64 `uri:"user_id"`
}

type BalanceReplenishment struct {
	UserID        uint64  `json:"user_id"`
	Replenishment float32 `json:"replenishment"`
}

type Balance struct {
	Balance float32 `json:"balance"`
}

// UpBalance godoc
// @Summary Пополнение или инициализация баланса
// @Schemes
// @Description Пополняет баланс пользователя или создаёт его при первом пополнении
// @Tags balance
// @Accept json
// @Param data body BalanceReplenishment true "Входные параметры"
// @Success 200
// @Router /balance/up_balance [post]
func (api *Api) UpBalance(ctx *gin.Context) {
	var replenishment BalanceReplenishment

	if err := ctx.BindJSON(&replenishment); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err := api.balance.UpsetUserBalance(ctx, models.UserBalance{UserID: replenishment.UserID, Balance: replenishment.Replenishment})
	if err == balance_service.ErrInvalidBalance {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	api.userTransaction.UpBalanceTransaction(ctx, replenishment.UserID, replenishment.Replenishment)

	ctx.AbortWithStatus(http.StatusOK)
}

// GetBalance godoc
// @Summary Получение баланса пользователя
// @Schemes
// @Description Возвращает баланс пользователя по его id
// @Tags balance
// @Accept json
// @Param data body UserID true "Входные параметры"
// @Produce json
// @Success 200 {object} Balance
// @Router /balance/get_balance [get]
func (api *Api) GetBalance(ctx *gin.Context) {
	var userID UserID

	if err := ctx.BindJSON(&userID); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	balance, err := api.balance.GetUserBalance(ctx, userID.ID)

	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, Balance{Balance: balance})
}
