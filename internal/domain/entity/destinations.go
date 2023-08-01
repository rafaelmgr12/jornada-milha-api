package entity

import (
	"errors"

	"github.com/google/uuid"
)

type Destinations struct {
	ID              uuid.UUID
	Name            string
	Price           float64
	Photo1          string
	Photo2          string
	Meta            string
	DescriptiveText string
}

func NewDestinations(name string, price float64, photo1 string, photo2 string, meta string, descriptiveText string) (*Destinations, error) {
	d := &Destinations{
		ID:              uuid.New(),
		Name:            name,
		Price:           price,
		Photo1:          photo1,
		Photo2:          photo2,
		Meta:            meta,
		DescriptiveText: descriptiveText,
	}
	err := d.Validate()
	if err != nil {
		return nil, err
	}
	return d, nil
}

func (d *Destinations) Validate() error {
	if d.Name == "" {
		return errors.New("name is required")
	}
	if d.Price < 0 {
		return errors.New("price is required")
	}
	if d.Photo1 == "" {
		return errors.New("photo1 is required")
	}
	if d.Photo2 == "" {
		return errors.New("photo2 is required")
	}
	if len(d.Meta) > 160 {
		return errors.New("meta cannot exceed 160 characters")
	}
	return nil
}
