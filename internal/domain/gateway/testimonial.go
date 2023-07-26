package gateway

import (
	"context"

	"github.com/rafaelmgr12/jornada-milha-api/internal/domain/entity"
)

type TestimonialGateway interface {
	CreateTestimonial(ctx context.Context, testimonial entity.Testimonial) (entity.Testimonial, error)
	ReadTestimonial(ctx context.Context) ([]entity.Testimonial, error)
	UpdateTestimonial(ctx context.Context, testimonial entity.Testimonial) (entity.Testimonial, error)
	DeleteTestimonial(ctx context.Context, id string) error
}
