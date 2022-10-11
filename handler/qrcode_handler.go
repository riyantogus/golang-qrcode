package handler

import (
	"golang-qrcode/controller"

	"github.com/labstack/echo/v4"
)

type QRCodeHandler interface {
	Handler(c *echo.Echo)
}

type qrCodeHandler struct{}

func NewQRCodeHandler() QRCodeHandler {
	return &qrCodeHandler{}
}

func (h *qrCodeHandler) Handler(c *echo.Echo) {
	qrCodeController := controller.NewQRCodeController()

	c.GET("/qrcode", qrCodeController.GenerateQRCode)
}
