package integration

import (
	"context"
	"testing"

	"github.com/rafaelmgr12/jornada-milha-api/internal/domain/entity"
	"github.com/rafaelmgr12/jornada-milha-api/internal/usecase/destinations"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockDestinationsGateway struct {
	mock.Mock
}

// GetDestinationsByID implements gateway.DestinationsGateway.
func (*mockDestinationsGateway) GetDestinationsByID(ctx context.Context, id string) (entity.Destinations, error) {
	panic("unimplemented")
}

// CreateDestinations implements gateway.DestinationsGateway.
func (*mockDestinationsGateway) CreateDestinations(ctx context.Context, destinations entity.Destinations) (entity.Destinations, error) {
	panic("unimplemented")
}

// DeleteDestinations implements gateway.DestinationsGateway.
func (*mockDestinationsGateway) DeleteDestinations(ctx context.Context, id string) error {
	panic("unimplemented")
}

// ReadDestinations implements gateway.DestinationsGateway.
func (*mockDestinationsGateway) ReadDestinations(ctx context.Context) ([]entity.Destinations, error) {
	panic("unimplemented")
}

// UpdateDestinations implements gateway.DestinationsGateway.
func (*mockDestinationsGateway) UpdateDestinations(ctx context.Context, destinations entity.Destinations) (entity.Destinations, error) {
	panic("unimplemented")
}

func (m *mockDestinationsGateway) GetDestinationsByName(ctx context.Context, name string) ([]entity.Destinations, error) {
	args := m.Called(ctx, name)
	return args.Get(0).([]entity.Destinations), args.Error(1)
}

func TestSearchDestinationsByName(t *testing.T) {
	// Mock do gateway
	gateway := &mockDestinationsGateway{}
	gateway.On("GetDestinationsByName", mock.Anything, "Paris").Return([]entity.Destinations{
		{
			Name:            "Paris",
			Price:           100,
			Photo1:          "photo_url",
			Photo2:          "photo_url",
			Meta:            "Teste",
			DescriptiveText: "PARIS TESTE",
		},
	}, nil)

	usecase := destinations.NewDestinationsUseCase(gateway)

	destinations, err := usecase.SearchDestinationsByName(context.Background(), "Paris")

	assert.NoError(t, err)
	assert.Len(t, destinations, 1)
	assert.Equal(t, "Paris", destinations[0].Name)
}
