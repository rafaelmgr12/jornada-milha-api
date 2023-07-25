package repository

import (
	"context"
	"database/sql"

	"github.com/rafaelmgr12/jornada-milha-api/internal/domain/entity"
	"github.com/rafaelmgr12/jornada-milha-api/internal/infra/db"
)

type TestimonialRepository struct {
	DB      *sql.DB
	Queries *db.Queries
}

func NewTestimonialRepository(dbt *sql.DB) *TestimonialRepository {
	return &TestimonialRepository{
		DB:      dbt,
		Queries: db.New(dbt),
	}
}

func (t *TestimonialRepository) CreateTestimonial(ctx context.Context, testimonial entity.Testimonial) error {
	err := t.Queries.CreateTestimonial(ctx, db.CreateTestimonialParams{
		ID:          testimonial.ID.String(),
		Name:        testimonial.Name,
		Testimonial: testimonial.Testimonial,
	})
	if err != nil {
		return err
	}
	return nil
}

func (t *TestimonialRepository) GetTestimonials(ctx context.Context) ([]entity.Testimonial, error) {
	res, err := t.Queries.GetTestimonial(ctx)
	if err != nil {
		return nil, err
	}
	testimonials := make([]entity.Testimonial, len(res))

	return testimonials, nil
}

func (t *TestimonialRepository) UpdateTestimonial(ctx context.Context, testimonial entity.Testimonial) error {
	err := t.Queries.UpdateTestimonial(ctx, db.UpdateTestimonialParams{
		ID:          testimonial.ID.String(),
		Name:        testimonial.Name,
		Testimonial: testimonial.Testimonial,
	})
	if err != nil {
		return err
	}
	return nil
}

func (t *TestimonialRepository) DeleteTestimonial(ctx context.Context, id string) error {
	err := t.Queries.DeleteTestimonial(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
