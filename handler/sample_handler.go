package handler

import (
	"golang-qrcode/controller"
	"golang-qrcode/util"

	"github.com/labstack/echo/v4"
)

type SampleHandler interface {
	Handler(c *echo.Echo)
}

type sampleHandler struct{}

func NewSampleHandler() QRCodeHandler {
	return &sampleHandler{}
}

func (h *sampleHandler) Handler(c *echo.Echo) {
	sampleController := controller.NewSampleController()
	renderer := util.NewTemplateRender()

	c.Renderer = renderer
	c.GET("/index", sampleController.Index)
}
