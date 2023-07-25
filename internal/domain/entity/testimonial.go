package entity

import (
	"errors"

	"github.com/google/uuid"
)

type Testimonial struct {
	ID          uuid.UUID
	Name        string
	Testimonial string
}

func NewTestimonial(name string, testimonial string) *Testimonial {
	return &Testimonial{
		ID:          uuid.New(),
		Name:        name,
		Testimonial: testimonial,
	}
}

func (t *Testimonial) Validate() error {
	if t.Name == "" {
		return errors.New("name is required")
	}
	if t.Testimonial == "" {
		return errors.New("testimonial is required")
	}
	return nil

}
