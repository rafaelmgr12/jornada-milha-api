package web

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rafaelmgr12/jornada-milha-api/internal/usecase/destinations"
)

type WebDestinationsHandler struct {
	DestinationsUseCase destinations.DestinationsUseCase
}

func NewWebDestinationsHandler(destinationsUseCase destinations.DestinationsUseCase) *WebDestinationsHandler {
	return &WebDestinationsHandler{
		DestinationsUseCase: destinationsUseCase,
	}
}

// CreateDestinations godoc
// @Summary Create a new destination
// @Description Create a new destination
// @Tags destinations
// @Accept json
// @Produce json
// @Param body body destinations.DestinationsCreateDTO  true        "DestinationsCreateDTO"
// @Success      200  {object}  entity.Destinations
// @Router       /api/v1/destinos [post]
func (h *WebDestinationsHandler) CreateDestinations(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if !json.Valid(body) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var dto destinations.DestinationsCreateDTO
	err = json.Unmarshal(body, &dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	createdDestinations, err := h.DestinationsUseCase.CreateDestinations(r.Context(), dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result, err := json.Marshal(createdDestinations)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
}

// GetListDestinations godoc
// @Summary Get list of destinations
// @Description Get list of destinations
// @Tags destinations
// @Accept json
// @Produce json
// @Success      200  {array}  entity.Destinations
// @Router       /api/v1/destinos [get]
func (h *WebDestinationsHandler) GetListDestinations(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	listDestinations, err := h.DestinationsUseCase.GetListDestinations(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result, err := json.Marshal(listDestinations)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
}

// UpdateDestinations godoc
// @Summary Update a destination
// @Description Update a destination
// @Tags destinations
// @Accept json
// @Produce json
// @Param  body body destinations.DestinationsUpdateDTO true 	  "DestinationUpdateDTO"
// Success      200  {object}  entity.Destinations
// @Router       /api/v1/destinos [put]
func (h *WebDestinationsHandler) UpdateDestinations(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if !json.Valid(body) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var dto destinations.DestinationsUpdateDTO
	err = json.Unmarshal(body, &dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	updatedDestinations, err := h.DestinationsUseCase.UpdateDestinations(r.Context(), dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result, err := json.Marshal(updatedDestinations)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
}

// DeleteDestinations godoc
// @Summary Delete a destination
// @Description Delete a destination
// @Tags destinations
// @Accept json
// @Produce json
// @Param  id path string true "Destination ID"
// @Router       /api/v1/destinos/{id} [delete]
func (h *WebDestinationsHandler) DeleteDestinations(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := h.DestinationsUseCase.DeleteDestinations(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// SearchDestinationsByName godoc
// @Summary Search destinations by name
// @Description Search destinations by name
// @Tags destinations
// @Accept json
// @Produce json
// @Param nome query string true "Destination Name"
// @Success 200 {array} entity.Destinations
// @Failure 404 {object} map[string]string "mensagem": "Nenhum destino foi encontrado"
// @Router /api/v1/query/destinos [get]
func (h *WebDestinationsHandler) SearchDestinationsByName(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("nome")
	if name == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	destinations, err := h.DestinationsUseCase.SearchDestinationsByName(r.Context(), name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if len(destinations) == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"mensagem": "Nenhum destino foi encontrado"})
		return
	}

	result, err := json.Marshal(destinations)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
}
