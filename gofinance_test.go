package gofinance

import "testing"

const key = ""

func TestGetQuotations(t *testing.T) {
	service := NewService("KEY")

	_, err := service.GetQuotations()
	if err != nil {
		t.Error(err)
		return
	}
}

func TestGetTaxes(t *testing.T) {
	service := NewService("KEY")

	_, err := service.GetTaxes()
	if err != nil {
		t.Error(err)
		return
	}
}
