package web

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rafaelmgr12/jornada-milha-api/internal/usecase/testimonials"
)

type WebTestimonialHandler struct {
	TestimonialUseCase testimonials.TestimonialsUseCase
}

func NewWebTestimonialHandler(testimonialUseCase testimonials.TestimonialsUseCase) *WebTestimonialHandler {
	return &WebTestimonialHandler{
		TestimonialUseCase: testimonialUseCase,
	}
}

// CreateTestimonial godoc
// @Summary      Create a Testimonial
// @Description  Create a Testimonial
// @Tags         testimonials
// @Accept       json
// @Produce      json
// @Param        body     body    testimonials.TestimonialCreateDTO   true        "TestimonialCreateDTO"
// @Success      200  {object}  entity.Testimonial
// @Router       /api/v1/depoimentos [post]
func (h *WebTestimonialHandler) CreateTestimonial(w http.ResponseWriter, r *http.Request) {

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

	var dto testimonials.TestimonialCreateDTO
	err = json.Unmarshal(body, &dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	createdTestimonial, err := h.TestimonialUseCase.CreateTestimonial(r.Context(), dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result, err := json.Marshal(createdTestimonial)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
}

// GetListTestimonials godoc
// @Summary      Retrieve all Testimonials
// @Description  Get a list of all Testimonials
// @Tags         testimonials
// @Accept       json
// @Produce      json
// @Success      200  {array}  entity.Testimonial
// @Router       /api/v1/depoimentos [get]
func (h *WebTestimonialHandler) GetListTestimonials(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	testimonials, err := h.TestimonialUseCase.GetListTestimonials(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result, err := json.Marshal(testimonials)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
}

// UpdateTestimonial godoc
// @Summary      Update a Testimonial
// @Description  Update a specific Testimonial
// @Tags         testimonials
// @Accept       json
// @Produce      json
// @Param        body     body    testimonials.TestimonialUpdateDTO   true        "TestimonialUpdateDTO"
// @Success      200  {object}  entity.Testimonial
// @Router       /api/v1/depoimentos [put]
func (h *WebTestimonialHandler) UpdateTestimonial(w http.ResponseWriter, r *http.Request) {
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

	var dto testimonials.TestimonialUpdateDTO
	err = json.Unmarshal(body, &dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res, err := h.TestimonialUseCase.UpdateTestimonial(r.Context(), dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
}

// DeleteTestimonial godoc
// @Summary      Delete a Testimonial
// @Description  Delete a specific Testimonial
// @Tags         testimonials
// @Accept       json
// @Produce      json
// @Param        id     path    string   true        "Testimonial ID"
// @Success      200
// @Router       /api/v1/depoimentos/{id} [delete]
func (h *WebTestimonialHandler) DeleteTestimonial(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := h.TestimonialUseCase.DeleteTestimonial(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// GetThreeRandonTestimonial godoc
// @Summary      Retrieve three random Testimonials
// @Description  Get a list of three random Testimonials
// @Tags         testimonials
// @Accept       json
// @Produce      json
// @Success      200  {array}  entity.Testimonial
// @Router       /api/v1/depoimentos-home [get]
func (h *WebTestimonialHandler) GetThreeRandonTestimonial(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	testimonials, err := h.TestimonialUseCase.GetThreeRandomTestimonials(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result, err := json.Marshal(testimonials)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
}
