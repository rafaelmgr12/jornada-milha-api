package gateway

import (
	"context"

	"github.com/rafaelmgr12/jornada-milha-api/internal/domain/entity"
)

type DestinationsGateway interface {
	CreateDestinations(ctx context.Context, destinations entity.Destinations) (entity.Destinations, error)
	ReadDestinations(ctx context.Context) ([]entity.Destinations, error)
	UpdateDestinations(ctx context.Context, destinations entity.Destinations) (entity.Destinations, error)
	DeleteDestinations(ctx context.Context, id string) error
	GetDestinationsByName(ctx context.Context, name string) ([]entity.Destinations, error)
	GetDestinationsByID(ctx context.Context, id string) (entity.Destinations, error)
}
