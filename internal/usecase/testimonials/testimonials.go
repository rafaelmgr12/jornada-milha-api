package testimonials

import (
	"context"

	"github.com/google/uuid"
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

type TestimonialListDTO struct {
	Testimonials []TestimonialListItemDTO `json:"testimonials"`
}

type TestimonialListItemDTO struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Testimonial string `json:"testimonial"`
}

func NewTestimonialsUseCase(testimonialGateway gateway.TestimonialGateway) *TestimonialsUseCase {
	return &TestimonialsUseCase{
		TestimonialGateway: testimonialGateway,
	}
}

func (uc *TestimonialsUseCase) CreateTestimonial(ctx context.Context, dto TestimonialCreateDTO) (entity.Testimonial, error) {
	testimonial := entity.Testimonial{
		ID:          uuid.New(),
		Name:        dto.Name,
		Testimonial: dto.Testimonial,
	}
	return uc.TestimonialGateway.CreateTestimonial(ctx, testimonial)

}

func (uc *TestimonialsUseCase) GetListTestimonials(ctx context.Context) (TestimonialListDTO, error) {
	listTestimonial, err := uc.TestimonialGateway.ReadTestimonial(ctx)
	if err != nil {
		return TestimonialListDTO{}, err
	}

	testimonialList := make([]TestimonialListItemDTO, len(listTestimonial))
	for i, testimonial := range listTestimonial {
		testimonialList[i] = TestimonialListItemDTO{
			ID:          testimonial.ID.String(),
			Name:        testimonial.Name,
			Testimonial: testimonial.Testimonial,
		}
	}

	return TestimonialListDTO{Testimonials: testimonialList}, nil
}

func (uc *TestimonialsUseCase) UpdateTestimonial(ctx context.Context, dto TestimonialUpdateDTO) (entity.Testimonial, error) {
	newTestimonial := entity.Testimonial{
		ID:          uuid.MustParse(dto.ID),
		Name:        dto.Name,
		Testimonial: dto.Testimonial,
	}

	return uc.TestimonialGateway.UpdateTestimonial(ctx, newTestimonial)

}

func (uc *TestimonialsUseCase) DeleteTestimonial(ctx context.Context, id string) error {
	return uc.TestimonialGateway.DeleteTestimonial(ctx, id)

}
