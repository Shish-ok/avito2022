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
// @Param user_id path string true "id пользователя"
// @Param data body BalanceReplenishment true "Входные параметры"
// @Success 200
// @Router /balance/up_balance/{user_id} [post]
func (api *Api) UpBalance(ctx *gin.Context) {
	var userID UserID
	if err := ctx.ShouldBindUri(&userID); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	var balance Balance
	if err := ctx.BindJSON(&balance); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err := api.balance.UpsetUserBalance(ctx, models.UserBalance{UserID: userID.ID, Balance: balance.Balance})
	if err == balance_service.ErrInvalidBalance {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.AbortWithStatus(http.StatusOK)
}

// GetBalance godoc
// @Summary Получение баланса пользователя
// @Schemes
// @Description Возвращает баланс пользователя по его id
// @Tags balance
// @Produce json
// @Success 200 {object} Balance
// @Param user_id  path string  true  "id пользователя"
// @Router /balance/get_balance/{user_id} [get]
func (api *Api) GetBalance(ctx *gin.Context) {
	var userID UserID
	if err := ctx.ShouldBindUri(&userID); err != nil {
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
