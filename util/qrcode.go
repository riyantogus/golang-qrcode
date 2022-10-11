package util

import (
	"fmt"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

type QRCodeUtil interface {
	GenerateQRCode(strURL string) (barcode.Barcode, error)
}

type qrCodeUtil struct{}

func NewQRCodeUtil() QRCodeUtil {
	return &qrCodeUtil{}
}

func (u *qrCodeUtil) GenerateQRCode(strURL string) (barcode.Barcode, error) {
	code, err := qr.Encode(strURL, qr.L, qr.Auto)
	if err != nil {
		return nil, fmt.Errorf("Cannot generate QR code.")
	}

	// Scale the barcode to the appropriate size
	code, err = barcode.Scale(code, 200, 200)
	if err != nil {
		return nil, fmt.Errorf("Cannot scale the QR code.")
	}
	return code, nil
}
