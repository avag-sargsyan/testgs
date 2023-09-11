package handler

import (
	"encoding/json"
	"github.com/avag-sargsyan/testgs/internal/conf"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

type MockPackService struct {
	mockCalculatePacks func(order int) map[int]int
}

func (m *MockPackService) CalculatePacks(order int) map[int]int {
	return m.mockCalculatePacks(order)
}

func TestPackHandler_Pack(t *testing.T) {
	// Create a mock PackService
	mockService := new(MockPackService)
	mockService.mockCalculatePacks = func(order int) map[int]int {
		return map[int]int{500: 1}
	}

	// Assuming that using default config pack sizes (250, 500, 1000, 2000, 5000)
	appConf := &conf.App{
		PackSizes: []int{250, 500, 1000, 2000, 5000},
	}
	h := NewPackHandler(appConf)
	h.(*packHandler).packService = mockService

	// Request with order=251
	req, err := http.NewRequest("GET", "/someRoute?order=251", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Handle response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.Pack)

	handler.ServeHTTP(rr, req)

	// Status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Response should be a map[int]int{500: 1}
	expected := map[int]int{500: 1}
	var responseMap map[int]int
	err = json.Unmarshal(rr.Body.Bytes(), &responseMap)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(responseMap, expected) {
		t.Errorf("handler returned unexpected body: got %v want %v", responseMap, expected)
	}
}
