package controller

import (
	"bytes"
	"golang-qrcode/helper"
	"golang-qrcode/util"
	"image/png"
	"net/http"
	"net/url"
	"strconv"

	"github.com/labstack/echo/v4"
)

type QRCodeController interface {
	GenerateQRCode(c echo.Context) error
}

type qrCodeController struct {
	QRCode util.QRCodeUtil
}

func NewQRCodeController() QRCodeController {
	return &qrCodeController{
		QRCode: util.NewQRCodeUtil(),
	}
}

func (q *qrCodeController) GenerateQRCode(c echo.Context) error {
	data := c.QueryParam("data")
	if data == "" {
		return c.JSON(http.StatusUnprocessableEntity, helper.ErrorResponse{
			Status:  http.StatusUnprocessableEntity,
			Message: "Something went wrong. Please try again.",
			Errors:  "Cannot get data.",
		})
	}

	strURL, err := url.QueryUnescape(data)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, helper.ErrorResponse{
			Status:  http.StatusUnprocessableEntity,
			Message: "Something went wrong. Please try again.",
			Errors:  "Cannot get query string.",
		})
	}

	code, err := q.QRCode.GenerateQRCode(strURL)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, helper.ErrorResponse{
			Status:  http.StatusUnprocessableEntity,
			Message: "Something went wrong. Please try again.",
			Errors:  err.Error(),
		})
	}

	buffer := new(bytes.Buffer)
	if err := png.Encode(buffer, code); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, helper.ErrorResponse{
			Status:  http.StatusUnprocessableEntity,
			Message: "Something went wrong. Please try again.",
			Errors:  "Cannoot encode png.",
		})
	}

	c.Response().Header().Set("Content-Type", "image/png")
	c.Response().Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))

	if _, err := c.Response().Write(buffer.Bytes()); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, helper.ErrorResponse{
			Status:  http.StatusUnprocessableEntity,
			Message: "Something went wrong. Please try again.",
			Errors:  "Cannot generate QR code png.",
		})
	}
	return nil
}
