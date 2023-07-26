package integration

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/rafaelmgr12/jornada-milha-api/internal/domain/entity"
	"github.com/stretchr/testify/assert"
)

func TestCreateTestimonial(t *testing.T) {
	// Create the mock
	mockRepo := &TestimonialRepositoryMock{
		CreateTestimonialFunc: func(ctx context.Context, testimonial entity.Testimonial) (entity.Testimonial, error) {
			// Define the behavior for the CreateTestimonial method
			return testimonial, nil
		},
	}

	// Use the mock in your test
	testimonial := entity.Testimonial{
		ID:          uuid.New(),
		Name:        "Test Name",
		Testimonial: "Test Testimonial",
	}
	result, err := mockRepo.CreateTestimonial(context.Background(), testimonial)

	// Check the results
	assert.Nil(t, err)
	assert.Equal(t, testimonial, result)
}
