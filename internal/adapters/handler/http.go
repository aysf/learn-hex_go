package handler

import (
	"gohex1/internal/core/domain"
	"gohex1/internal/core/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HTTPHandler struct {
	svc services.MessengerService
}

func NewHTTPHandler(messangerService services.MessengerService) *HTTPHandler {
	return &HTTPHandler{
		svc: messangerService,
	}
}

func (h *HTTPHandler) SaveMessage(ctx *gin.Context) {
	var message domain.Message

	if err := ctx.ShouldBindJSON(&message); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	err := h.svc.SaveMessage(message)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "New message created successfully",
	})

}

func (h *HTTPHandler) ReadMessage(ctx *gin.Context) {
	id := ctx.Param("id")
	message, err := h.svc.ReadMessage(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, message)
}

func (h *HTTPHandler) ReadMessages(ctx *gin.Context) {

	messages, err := h.svc.ReadMessages()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, messages)
}
