package usecase

import (
	"sort"
	"sync"

	"github.com/avag-sargsyan/testgs/internal/conf"
)

type Pack struct {
	Size int
}

var (
	once      sync.Once
	singleton PackService
)

type PackService interface {
	CalculatePacks(order int) map[int]int
}

type packService struct {
	packSizes []Pack
}

func NewPackService(conf *conf.App) PackService {
	var packSizes []Pack
	for _, size := range conf.PackSizes {
		packSizes = append(packSizes, Pack{Size: size})
	}
	return &packService{packSizes: packSizes}
}

func GetPackService(conf *conf.App) PackService {
	once.Do(func() {
		singleton = NewPackService(conf)
	})
	return singleton
}

func (s *packService) CalculatePacks(order int) map[int]int {
	result := make(map[int]int)

	// Sort packSizes in descending order to start with the biggest pack size
	sort.Slice(s.packSizes, func(i, j int) bool {
		return s.packSizes[i].Size > s.packSizes[j].Size
	})

	for i := 0; i < len(s.packSizes); i++ {
		// If order is bigger than pack size, calculate the number of packs
		if order >= s.packSizes[i].Size {
			numPacks := order / s.packSizes[i].Size
			result[s.packSizes[i].Size] = numPacks
			order = order % s.packSizes[i].Size
		} else {
			// If order is less than the smallest pack size, add the smallest pack size to the result
			min := s.packSizes[len(s.packSizes)-1]
			if s.packSizes[i].Size-order < min.Size {
				result[s.packSizes[i].Size] = 1
				order = 0
				continue
			}

		}
	}

	return result
}
