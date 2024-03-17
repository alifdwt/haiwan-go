package rajaongkir

import (
	"net/http"

	"github.com/alifdwt/haiwan-go/pkg/dotenv"
	"github.com/gofiber/fiber/v2/log"
)

type RajaOngkirAPI struct {
	BaseURL string
	ApiKey  string
}

func NewRajaOngkirAPI() *RajaOngkirAPI {
	config, err := dotenv.LoadConfig("../..")
	if err != nil {
		log.Error("cannot load .env file", err)
	}

	return &RajaOngkirAPI{
		BaseURL: "https://api.rajaongkir.com",
		ApiKey:  config.RAJAONGKIR_API,
	}
}

func (api *RajaOngkirAPI) GetConnectionAndHeaders() (*http.Client, map[string]string) {
	client := &http.Client{}
	headers := map[string]string{
		"key":          api.ApiKey,
		"Content-Type": "application/x-www-form-urlencoded",
	}

	return client, headers
}
