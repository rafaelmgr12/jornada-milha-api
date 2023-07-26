package web

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

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

func (h *WebTestimonialHandler) DeleteTestimonial(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")
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
