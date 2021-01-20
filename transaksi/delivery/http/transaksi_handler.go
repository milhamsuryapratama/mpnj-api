package http

import (
	"context"
	"github.com/gin-gonic/gin"
	"mpnj-api/domain"
	"mpnj-api/utils"
	"strconv"
)

type TransaksiHandler struct {
	TransaksiUsecase domain.TransaksiUsecase
}

func TransaksiHandlerFunc(r *gin.RouterGroup, transaksi domain.TransaksiUsecase)  {
	handler := &TransaksiHandler{
		TransaksiUsecase: transaksi,
	}

	r.POST("/transaksi", handler.CreateTransaksi)
	r.GET("/transaksi", handler.GetTransaksi)
}

func (t *TransaksiHandler) CreateTransaksi(c *gin.Context) {
	var (
		transaksi domain.Transaksi
	)
	err := c.ShouldBindJSON(&transaksi)
	if err != nil {
		utils.Render(c, err.Error(), nil, 400)
		return
	}

	err = t.TransaksiUsecase.Create(c, &transaksi)
	if err != nil {
		utils.Render(c, err.Error(), nil, 400)
		return
	}

	utils.Render(c, "sukses", transaksi, 201)
}

func (t *TransaksiHandler) GetTransaksi(c *gin.Context) {
	var (
		ctx context.Context
		param, _ = strconv.Atoi(c.Query("page"))
	)

	transaksi, err := t.TransaksiUsecase.Get(ctx, param)

	if err != nil {
		utils.Render(c, err.Error(), nil, 400)
		return
	}

	utils.Render(c, "sukses", transaksi, 200)
}