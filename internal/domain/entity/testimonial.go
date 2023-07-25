package entity

import (
	"errors"

	"github.com/google/uuid"
)

type Testimonial struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Testimonial string    `json:"testimonial"`
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
