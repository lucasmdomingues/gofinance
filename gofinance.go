package gofinance

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

const URL = "https://api.hgbrasil.com/finance"

type service struct {
	Key string
}

type Service interface {
	GetQuotations() (*DataQuotation, error)
	GetTaxes() (*DataTaxe, error)
}

func NewService(key string) Service {
	return &service{
		Key: key,
	}
}

func (s *service) GetQuotations() (*DataQuotation, error) {
	url := fmt.Sprintf("%s/quotations?key=%s", URL, s.Key)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("Oops, error on get quotations")
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var dataQuotation *DataQuotation

	err = json.Unmarshal(body, &dataQuotation)
	if err != nil {
		return nil, err
	}

	return dataQuotation, nil
}

func (s *service) GetTaxes() (*DataTaxe, error) {
	url := fmt.Sprintf("%s/taxes?key=%s", URL, s.Key)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("Oops, error on get taxes")
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var dataTaxe *DataTaxe

	err = json.Unmarshal(body, &dataTaxe)
	if err != nil {
		return nil, err
	}

	return dataTaxe, nil
}
