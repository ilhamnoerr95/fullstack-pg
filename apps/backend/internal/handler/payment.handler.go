package handler

import (
	"backend/internal/service"

	"github.com/gin-gonic/gin"
)	

type PaymentHandler struct {
	service service.PaymentService
}

func NewPaymentHandler(s service.PaymentService) *PaymentHandler {
	return &PaymentHandler{service: s}
}


func (h *PaymentHandler) GetAll(c *gin.Context) {

	status := c.Query("status")

	if status != "" {
		data, err := h.service.GetPaymentsByStatus(status)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"payments": data})
		return
	}

	data, summary, err := h.service.GetAll()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"payments": data,
		"summary":  summary,
	})
}

