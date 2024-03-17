package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	rajaongkirrequest "github.com/alifdwt/haiwan-go/internal/domain/requests/rajaongkir_request"
	"github.com/alifdwt/haiwan-go/pkg/rajaongkir"
)

type rajaOngkirService struct {
	rajaOngkir *rajaongkir.RajaOngkirAPI
}

func NewRajaOngkirService(rajaOngkir *rajaongkir.RajaOngkirAPI) *rajaOngkirService {
	return &rajaOngkirService{rajaOngkir: rajaOngkir}
}

func (r *rajaOngkirService) GetProvince() (map[string]interface{}, error) {
	endpoint := "/starter/province"
	client, headers := r.rajaOngkir.GetConnectionAndHeaders()
	url := fmt.Sprintf("%s%s", r.rajaOngkir.BaseURL, endpoint)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusOK {
		var result map[string]interface{}

		err := json.NewDecoder(res.Body).Decode(&result)
		if err != nil {
			return nil, err
		}

		return result, nil
	}

	return nil, errors.New("failed to fetch province data from RajaOngkir API. Status code: " + string(rune(res.StatusCode)))
}

func (r *rajaOngkirService) GetCity(provinceId int) (map[string]interface{}, error) {
	endpoint := fmt.Sprintf("/starter/city?province=%d", provinceId)
	client, headers := r.rajaOngkir.GetConnectionAndHeaders()
	url := fmt.Sprintf("%s%s", r.rajaOngkir.BaseURL, endpoint)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusOK {
		var result map[string]interface{}

		err := json.NewDecoder(res.Body).Decode(&result)
		if err != nil {
			return nil, err
		}

		return result, nil
	}

	return nil, errors.New("Failed to fetch city data from RajaOngkir API. Status code: " + string(rune(res.StatusCode)))
}

func (r *rajaOngkirService) GetCost(request rajaongkirrequest.OngkosRequest) (map[string]interface{}, error) {
	endpoint := "/starter/cost"
	client, headers := r.rajaOngkir.GetConnectionAndHeaders()
	url := fmt.Sprintf("%s%s", r.rajaOngkir.BaseURL, endpoint)

	payload := fmt.Sprintf("origin=%s&destination=%s&weight%d&courier=%s",
		request.Asal, request.Tujuan, request.Berat, request.Kurir)

	req, err := http.NewRequest("POST", url, strings.NewReader(payload))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusOK {
		var result map[string]interface{}

		err := json.NewDecoder(res.Body).Decode(&result)
		if err != nil {
			return nil, err
		}

		return result, nil
	}

	return nil, errors.New("Failed to get shipping cost from RajaOngkir API. Status code: " + string(rune(res.StatusCode)))
}
