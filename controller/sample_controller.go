package controller

import (
	"fmt"
	"golang-qrcode/dto"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/thanhpk/randstr"
)

type SampleController interface {
	Index(c echo.Context) error
}

type sampleController struct{}

func NewSampleController() SampleController {
	return &sampleController{}
}

func (s *sampleController) Index(c echo.Context) error {
	//baseURL
	baseURL := os.Getenv("BASE_URL")

	// Generate QrCode
	qrcode := randstr.String(6)

	data := dto.Sample{
		QRCodeURL: fmt.Sprintf("%s/qrcode?data=%s", baseURL, qrcode),
	}

	return c.Render(http.StatusOK, "qrcode.html", data)
}
