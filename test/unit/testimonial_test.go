package unit

import (
	"testing"

	"github.com/rafaelmgr12/jornada-milha-api/internal/domain/entity"
	"github.com/stretchr/testify/assert"
)

func TestTestimonial_Validate(t *testing.T) {
	testimonial := entity.NewTestimonial("", "")

	err := testimonial.Validate()

	assert.NotNil(t, err)
	assert.Equal(t, "name is required", err.Error())

	testimonial.Name = "Test Name"
	err = testimonial.Validate()

	assert.NotNil(t, err)
	assert.Equal(t, "testimonial is required", err.Error())

	testimonial.Testimonial = "Test Testimonial"
	err = testimonial.Validate()

	assert.Nil(t, err)
}
