package api

import (
	"avito2022/internal/app/models"
	"avito2022/internal/app/service/balance_holder"
	"github.com/gin-gonic/gin"
	"net/http"
)

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
// @Router /sales/reserve_money [post]
func (api *Api) ReserveMoney(ctx *gin.Context) {
	var reserve Reserve

	if err := ctx.BindJSON(&reserve); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err := api.balanceHolder.ReserveMoney(ctx, models.HolderOperation{
		OrderID:     reserve.OrderID,
		ServiceID:   reserve.ServiceID,
		UserID:      reserve.UserID,
		ServiceName: reserve.ServiceName,
		Cost:        reserve.Cost,
	})

	if err == balance_holder.ErrInvalidCost {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	if err == balance_holder.ErrInsufficientBalance {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.AbortWithStatus(http.StatusOK)
}
