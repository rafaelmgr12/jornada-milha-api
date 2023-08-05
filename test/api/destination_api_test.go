package api

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/rafaelmgr12/jornada-milha-api/internal/domain/entity"
	"github.com/rafaelmgr12/jornada-milha-api/internal/infra/web"
	"github.com/rafaelmgr12/jornada-milha-api/internal/usecase/chat"
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
func (m *mockDestinationsGateway) CreateDestinations(ctx context.Context, destinations entity.Destinations) (entity.Destinations, error) {
	args := m.Called(ctx, destinations)
	return args.Get(0).(entity.Destinations), args.Error(1)
}

func (m *mockDestinationsGateway) DeleteDestinations(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

// ReadDestinations implements gateway.DestinationsGateway.
func (m *mockDestinationsGateway) ReadDestinations(ctx context.Context) ([]entity.Destinations, error) {
	args := m.Called(ctx)
	return args.Get(0).([]entity.Destinations), args.Error(1)
}

func (m *mockDestinationsGateway) UpdateDestinations(ctx context.Context, destinations entity.Destinations) (entity.Destinations, error) {
	args := m.Called(ctx, destinations)
	return args.Get(0).(entity.Destinations), args.Error(1)
}

func (m *mockDestinationsGateway) GetDestinationsByName(ctx context.Context, name string) ([]entity.Destinations, error) {
	args := m.Called(ctx, name)
	return args.Get(0).([]entity.Destinations), args.Error(1)
}

func TestGetListDestinations(t *testing.T) {
	expectedJson := `{"destinations":[{"id":"00000000-0000-0000-0000-000000000000","name":"Teste","price":10,"photo_1":"teste","photo_2":"teste","meta":"teste","descriptive_text":"teste"}]}`
	r := httptest.NewRequest(http.MethodGet, "/destinos", nil)
	w := httptest.NewRecorder()

	// Mock do gateway
	gateway := &mockDestinationsGateway{}
	gateway.On("ReadDestinations", mock.Anything).Return([]entity.Destinations{
		{
			ID:              uuid.MustParse("00000000-0000-0000-0000-000000000000"),
			Name:            "Teste",
			Price:           10,
			Photo1:          "teste",
			Photo2:          "teste",
			Meta:            "teste",
			DescriptiveText: "teste",
		},
	}, nil)

	createDescription := chat.CreateDescriptionUseCase{
		APIKey: "Chave-Teste",
	}

	usecase := destinations.NewDestinationsUseCase(gateway)

	handler := web.NewWebDestinationsHandler(*usecase, createDescription) // Passando o ponteiro para o caso de uso
	handler.GetListDestinations(w, r)

	resBody, err := io.ReadAll(w.Result().Body)
	if err != nil {
		t.Fatalf("unexpected while reading body: %v", err)
	}
	w.Result().Body.Close()
	assert.Equal(t, expectedJson, string(resBody))
}

func TestCreateDestination(t *testing.T) {
	inputJson := `{"name":"Teste","price":10,"photo_1":"teste","photo_2":"teste","meta":"teste","descriptive_text":"teste"}`
	expectedJson := `{"ID":"00000000-0000-0000-0000-000000000000","Name":"Teste","Price":10,"Photo1":"teste","Photo2":"teste","Meta":"teste","DescriptiveText":"teste"}`

	r := httptest.NewRequest(http.MethodPost, "/destinos", strings.NewReader(inputJson))
	w := httptest.NewRecorder()

	gateway := &mockDestinationsGateway{}
	gateway.On("CreateDestinations", mock.Anything, mock.Anything).Return(entity.Destinations{
		ID:              uuid.MustParse("00000000-0000-0000-0000-000000000000"),
		Name:            "Teste",
		Price:           10,
		Photo1:          "teste",
		Photo2:          "teste",
		Meta:            "teste",
		DescriptiveText: "teste",
	}, nil)

	createDescription := chat.CreateDescriptionUseCase{
		APIKey: "Chave-Teste",
	}

	usecase := destinations.NewDestinationsUseCase(gateway)
	handler := web.NewWebDestinationsHandler(*usecase, createDescription)
	handler.CreateDestinations(w, r)

	resBody, err := io.ReadAll(w.Result().Body)
	if err != nil {
		t.Fatalf("unexpected while reading body: %v", err)
	}
	w.Result().Body.Close()
	assert.Equal(t, expectedJson, string(resBody))
}

func TestUpdateDestination(t *testing.T) {
	inputJson := `{"id":"00000000-0000-0000-0000-000000000000","name":"Teste Updated","price":20,"photo_1":"teste","photo_2":"teste","meta":"teste","descriptive_text":"teste_updated"}`
	r := httptest.NewRequest(http.MethodPut, "/destinos", strings.NewReader(inputJson))
	w := httptest.NewRecorder()

	gateway := &mockDestinationsGateway{}
	gateway.On("UpdateDestinations", mock.Anything, mock.Anything).Return(entity.Destinations{
		ID:              uuid.MustParse("00000000-0000-0000-0000-000000000000"),
		Name:            "Teste",
		Price:           20,
		Photo1:          "teste",
		Photo2:          "teste",
		Meta:            "teste",
		DescriptiveText: "teste_updated",
	}, nil)

	createDescription := chat.CreateDescriptionUseCase{
		APIKey: "Chave-Teste",
	}

	usecase := destinations.NewDestinationsUseCase(gateway)
	handler := web.NewWebDestinationsHandler(*usecase, createDescription)
	handler.UpdateDestinations(w, r)

	assert.Equal(t, http.StatusOK, w.Result().StatusCode)
}
