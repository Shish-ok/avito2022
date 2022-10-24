package api

import "github.com/gin-gonic/gin"

type Reserve struct {
	OrderID     uint64  `json:"order_id"`
	ServiceID   uint64  `json:"service_id"`
	UserID      uint64  `json:"user_id"`
	ServiceName string  `json:"service_name"`
	Cost        float32 `json:"cost"`
}

// ReserveMoney godoc
// @Summary Резервирование средств
// @Schemes
// @Description Резервирование средств с основного баланса на отдельном счете
// @Tags sales
// @Accept json
// @Param data body Reserve true "Входные параметры"
// @Success 200
// @Router /balance/reserve_money [post]
func (api *Api) ReserveMoney(ctx *gin.Context) {

}
