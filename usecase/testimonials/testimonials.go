package testimonials

import (
	"context"

	"github.com/rafaelmgr12/jornada-milha-api/internal/domain/entity"
	"github.com/rafaelmgr12/jornada-milha-api/internal/domain/gateway"
)

type TestimonialsUseCase struct {
	TestimonialGateway gateway.TestimonialGateway
}

type TestimonialCreateDTO struct {
	Name        string `json:"name"`
	Testimonial string `json:"Testimonial"`
}

type TestimonialUpdateDTO struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Testimonial string `json:"Testimonial"`
}

func NewTestimonialsUseCase(testimonialGateway gateway.TestimonialGateway) *TestimonialsUseCase {
	return &TestimonialsUseCase{
		TestimonialGateway: testimonialGateway,
	}
}

func (uc *TestimonialsUseCase) CreateTestimonial(ctx context.Context, dto TestimonialCreateDTO) error {
	return uc.TestimonialGateway.CreateTestimonial(ctx, dto.Name, dto.Testimonial)

}

func (uc *TestimonialsUseCase) GetListTestimonials(ctx context.Context) ([]*entity.Testimonial, error) {
	return uc.TestimonialGateway.ReadTestimonial(ctx)
}

func (uc *TestimonialsUseCase) UpdateTestimonial(ctx context.Context, dto TestimonialUpdateDTO) error {
	return uc.TestimonialGateway.UpdateTestimonial(ctx, dto.ID, dto.Name, dto.Testimonial)

}

func (uc *TestimonialsUseCase) DeleteTestimonial(ctx context.Context, id string) error {
	return uc.TestimonialGateway.DeleteTestimonial(ctx, id)

}
