package destinations

import (
	"context"

	"github.com/google/uuid"
	"github.com/rafaelmgr12/jornada-milha-api/internal/domain/entity"
	"github.com/rafaelmgr12/jornada-milha-api/internal/domain/gateway"
)

type DestinationsUseCase struct {
	DestinationsGateway gateway.DestinationsGateway
}

type DestinationsCreateDTO struct {
	Name            string  `json:"name"`
	Price           float64 `json:"price"`
	Photo1          string  `json:"photo_1"`
	Photo2          string  `json:"photo_2"`
	Meta            string  `json:"meta"`
	DescriptiveText string  `json:"descriptive_text"`
}

type DestinationsUpdateDTO struct {
	ID              string  `json:"id"`
	Name            string  `json:"name"`
	Price           float64 `json:"price"`
	Photo1          string  `json:"photo_1"`
	Photo2          string  `json:"photo_2"`
	Meta            string  `json:"meta"`
	DescriptiveText string  `json:"descriptive_text"`
}

type DestinationsListDTO struct {
	Destinations []DestinationsListItemDTO `json:"destinations"`
}

type DestinationsListItemDTO struct {
	ID              string  `json:"id"`
	Name            string  `json:"name"`
	Price           float64 `json:"price"`
	Photo1          string  `json:"photo_1"`
	Photo2          string  `json:"photo_2"`
	Meta            string  `json:"meta"`
	DescriptiveText string  `json:"descriptive_text"`
}

func NewDestinationsUseCase(destinationsGateway gateway.DestinationsGateway) *DestinationsUseCase {
	return &DestinationsUseCase{
		DestinationsGateway: destinationsGateway,
	}
}

func (uc *DestinationsUseCase) CreateDestinations(ctx context.Context, dto DestinationsCreateDTO) (entity.Destinations, error) {
	destinations := entity.Destinations{
		ID:              uuid.New(),
		Name:            dto.Name,
		Price:           dto.Price,
		Photo1:          dto.Photo1,
		Photo2:          dto.Photo2,
		Meta:            dto.Meta,
		DescriptiveText: dto.DescriptiveText,
	}
	return uc.DestinationsGateway.CreateDestinations(ctx, destinations)
}

func (uc *DestinationsUseCase) GetListDestinations(ctx context.Context) (DestinationsListDTO, error) {
	listDestinations, err := uc.DestinationsGateway.ReadDestinations(ctx)
	if err != nil {
		return DestinationsListDTO{}, err
	}

	destinationsList := make([]DestinationsListItemDTO, len(listDestinations))
	for i, destinations := range listDestinations {
		destinationsList[i] = DestinationsListItemDTO{
			ID:              destinations.ID.String(),
			Name:            destinations.Name,
			Price:           destinations.Price,
			Photo1:          destinations.Photo1,
			Photo2:          destinations.Photo2,
			Meta:            destinations.Meta,
			DescriptiveText: destinations.DescriptiveText,
		}
	}

	return DestinationsListDTO{Destinations: destinationsList}, nil
}

func (uc *DestinationsUseCase) UpdateDestinations(ctx context.Context, dto DestinationsUpdateDTO) (entity.Destinations, error) {
	newDestination := entity.Destinations{
		ID:              uuid.MustParse(dto.ID),
		Name:            dto.Name,
		Price:           dto.Price,
		Photo1:          dto.Photo1,
		Photo2:          dto.Photo2,
		Meta:            dto.Meta,
		DescriptiveText: dto.DescriptiveText,
	}

	return uc.DestinationsGateway.UpdateDestinations(ctx, newDestination)
}

func (uc *DestinationsUseCase) DeleteDestinations(ctx context.Context, id string) error {
	return uc.DestinationsGateway.DeleteDestinations(ctx, id)
}

func (uc *DestinationsUseCase) SearchDestinationsByName(ctx context.Context, name string) ([]entity.Destinations, error) {
	destinations, err := uc.DestinationsGateway.GetDestinationsByName(ctx, name)
	if err != nil {
		return nil, err
	}

	return destinations, nil
}
