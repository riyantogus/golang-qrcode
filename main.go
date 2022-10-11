package main

import (
	"encoding/json"
	"golang-qrcode/handler"
	"io/ioutil"
	"log"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func init() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Failed to load env file.")
	}
}

func main() {
	e := echo.New()

	qrCodeHandler := handler.NewQRCodeHandler()
	qrCodeHandler.Handler(e)

	sampleHandler := handler.NewSampleHandler()
	sampleHandler.Handler(e)

	// list of registered route path
	data, err := json.MarshalIndent(e.Routes(), "", " ")
	if err != nil {
		log.Fatal(err)
	}
	ioutil.WriteFile("./routes.json", data, 0644)

	e.Logger.Fatal(e.Start(":9090"))
}
