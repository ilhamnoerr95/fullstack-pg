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
	data, summary, _ := h.service.GetAll()

	c.JSON(200, gin.H{
		"payments": data,
		"summary":  summary,
	})
}