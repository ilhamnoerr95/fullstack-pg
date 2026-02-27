package handler

import (
	"net/http"

	"backend/internal/service"

	"github.com/gin-gonic/gin"
)

type PaymentHandler struct {
	service *service.PaymentService
}

func NewPaymentHandler(s *service.PaymentService) *PaymentHandler {
	return &PaymentHandler{service: s}
}

func (h *PaymentHandler) GetAll(c *gin.Context) {

	status := c.Query("status")

	//  Kalau ada status → filter
	if status != "" {
		data, err := h.service.GetPaymentsByStatus(status)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"payments": data,
		})
		return
	}

	// Kalau tidak ada → ambil semua (pakai Redis cache)
	data, summary, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"payments": data,
		"summary":  summary,
	})
}