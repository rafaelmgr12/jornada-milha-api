package repository

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/rafaelmgr12/jornada-milha-api/internal/domain/entity"
	"github.com/rafaelmgr12/jornada-milha-api/internal/infra/db"
)

type DestinationsRepository struct {
	DB      *sql.DB
	Queries *db.Queries
}

func NewDestinationsRepository(dbt *sql.DB) *DestinationsRepository {
	return &DestinationsRepository{
		DB:      dbt,
		Queries: db.New(dbt),
	}
}

func (d *DestinationsRepository) CreateDestinations(ctx context.Context, destinations entity.Destinations) (entity.Destinations, error) {

	err := d.Queries.CreateDestination(ctx, db.CreateDestinationParams{
		ID:    destinations.ID.String(),
		Name:  destinations.Name,
		Price: destinations.Price,
		Photo: destinations.Photo,
	})
	if err != nil {
		return entity.Destinations{}, err
	}

	return destinations, nil
}

func (d *DestinationsRepository) ReadDestinations(ctx context.Context) ([]entity.Destinations, error) {

	res, err := d.Queries.GetDestination(ctx)
	if err != nil {
		return nil, err
	}

	destinations := make([]entity.Destinations, len(res))

	for i, destination := range res {
		destinations[i] = entity.Destinations{
			ID:    uuid.MustParse(destination.ID),
			Name:  destination.Name,
			Price: destination.Price,
			Photo: destination.Photo,
		}
	}

	return destinations, nil

}

func (d *DestinationsRepository) UpdateDestinations(ctx context.Context, destinations entity.Destinations) (entity.Destinations, error) {
	err := d.Queries.UpdateDestination(ctx, db.UpdateDestinationParams{
		ID:    destinations.ID.String(),
		Name:  destinations.Name,
		Price: destinations.Price,
		Photo: destinations.Photo,
	})

	if err != nil {
		return entity.Destinations{}, err
	}

	return destinations, nil
}

func (d *DestinationsRepository) DeleteDestinations(ctx context.Context, id string) error {
	err := d.Queries.DeleteDestination(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (d *DestinationsRepository) GetDestinationsByName(ctx context.Context, name string) ([]entity.Destinations, error) {
	res, err := d.Queries.GetDestinationsByName(ctx, name)
	if err != nil {
		return nil, err
	}

	destinations := make([]entity.Destinations, len(res))
	for i, destination := range res {
		destinations[i] = entity.Destinations{
			ID:    uuid.MustParse(destination.ID),
			Name:  destination.Name,
			Price: destination.Price,
			Photo: destination.Photo,
		}
	}

	return destinations, nil
}
