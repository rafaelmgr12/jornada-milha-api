package gateway

import (
	"context"

	"github.com/rafaelmgr12/jornada-milha-api/internal/domain/entity"
)

type TestimonialGateway interface {
	CreateTestimonial(ctx context.Context, name string, testimonial string) error
	ReadTestimonial(ctx context.Context) ([]*entity.Testimonial, error)
	UpdateTestimonial(ctx context.Context, id string, name string, testimonial string) error
	DeleteTestimonial(ctx context.Context, id string) error
}
