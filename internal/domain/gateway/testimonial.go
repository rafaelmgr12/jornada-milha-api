package gateway

import "context"

type TestimonialGateway interface {
	CreateTestimonial(ctx context.Context, name string, testimonial string) error
	ReadTestimonial(ctx context.Context, id string) (string, string, error)
	UpdateTestimonial(ctx context.Context, id string, name string, testimonial string) error
	DeleteTestimonial(ctx context.Context, id string) error
}
