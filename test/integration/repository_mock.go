package integration

import (
	"context"

	"github.com/rafaelmgr12/jornada-milha-api/internal/domain/entity"
)

type TestimonialRepositoryMock struct {
	CreateTestimonialFunc func(ctx context.Context, testimonial entity.Testimonial) (entity.Testimonial, error)
	ReadTestimonialFunc   func(ctx context.Context) ([]entity.Testimonial, error)
	UpdateTestimonialFunc func(ctx context.Context, testimonial entity.Testimonial) (entity.Testimonial, error)
	DeleteTestimonialFunc func(ctx context.Context, id string) error
}

func (m *TestimonialRepositoryMock) CreateTestimonial(ctx context.Context, testimonial entity.Testimonial) (entity.Testimonial, error) {
	if m.CreateTestimonialFunc == nil {
		return entity.Testimonial{}, nil
	}
	return m.CreateTestimonialFunc(ctx, testimonial)
}
