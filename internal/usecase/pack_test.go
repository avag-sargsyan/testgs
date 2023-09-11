package usecase

import (
	"github.com/avag-sargsyan/testgs/internal/conf"
	"reflect"
	"testing"
)

func TestCalculatePacks(t *testing.T) {
	// Assuming that using default config pack sizes (250, 500, 1000, 2000, 5000)
	appConf := &conf.App{
		PackSizes: []int{250, 500, 1000, 2000, 5000},
	}

	packService := NewPackService(appConf)

	tests := []struct {
		order  int
		output map[int]int
	}{
		{0, map[int]int{}},
		{1, map[int]int{250: 1}},
		{251, map[int]int{500: 1}},
		{501, map[int]int{500: 1, 250: 1}},
		{10000, map[int]int{5000: 2}},
		{10001, map[int]int{5000: 2, 250: 1}},
		{12001, map[int]int{5000: 2, 2000: 1, 250: 1}},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			if got := packService.CalculatePacks(test.order); !reflect.DeepEqual(got, test.output) {
				t.Errorf("CalculatePacks(%d) = %v; want %v", test.order, got, test.output)
			}
		})
	}
}
