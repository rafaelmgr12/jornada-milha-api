package entity

import (
	"errors"

	"github.com/google/uuid"
)

type Destinations struct {
	ID    uuid.UUID
	Name  string
	Price float64
	Photo string
}

func NewDestinations(name string, price float64, photo string) *Destinations {
	return &Destinations{
		ID:    uuid.New(),
		Name:  name,
		Price: price,
		Photo: photo,
	}
}

func (d *Destinations) Validate() error {
	if d.Name == "" {
		return errors.New("name is required")
	}
	if d.Price < 0 {
		return errors.New("price is required")
	}
	if d.Photo == "" {
		return errors.New("photo is required")
	}
	return nil

}
