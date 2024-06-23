package domain

import (
	"errors"
	"fmt"
)

type SpotService struct{}

var (
	ErrInvalidQuantity = errors.New("quantity must be greater than zero")
)

func NewSpotService() *SpotService {
	return &SpotService{}
}

func (s *SpotService) GenerateSpots(event *Event, quantity int) error {
	if quantity <= 0 {
		return ErrInvalidQuantity
	}

	for i := range quantity {
		spotName := fmt.Sprintf("%c", 'A'+i+1)
		spot, err := NewSpot(event, spotName)
		if err != nil {
			return err
		}
		event.Spots = append(event.Spots, *spot)
	}

	return nil
}
