package api

import "github.com/gin-gonic/gin"

type Balance struct {
	Balance float32 `json:"balance"`
}

// UpBalance godoc
// @Summary Пополнение или инициализация баланса
// @Schemes
// @Description Пополняет баланс пользователя или создаёт его при первом пополнении
// @Tags balance
// @Accept json
// @Param data body AddBalance true "Входные параметры"
// @Success 200
// @Router /balance/up_balance/{user_id} [post]
func (api *Api) UpBalance(ctx *gin.Context) {

}

// GetBalance godoc
// @Summary Получение баланса пользователя
// @Schemes
// @Description Возвращает баланс пользователя по его id
// @Tags balance
// @Produce json
// @Success 200 {object} Balance
// @Router /balance/get_balance/{user_id} [get]
func (api *Api) GetBalance(ctx *gin.Context) {

}
