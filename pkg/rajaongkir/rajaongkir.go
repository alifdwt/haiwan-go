package rajaongkir

import (
	"net/http"

	"github.com/spf13/viper"
)

type RajaOngkirAPI struct {
	BaseURL string
	ApiKey  string
}

func NewRajaOngkirAPI() *RajaOngkirAPI {
	// config, err := dotenv.LoadConfig(".")
	// if err != nil {
	// 	log.Error("cannot load .env file", err)
	// }

	return &RajaOngkirAPI{
		BaseURL: "https://api.rajaongkir.com",
		ApiKey:  viper.GetString("RAJAONGKIR_API"),
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
