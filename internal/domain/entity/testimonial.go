package entity

import (
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
