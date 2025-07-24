package http

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Shyyw1e/effective-mobile-subs/internal/usecase"
	"github.com/Shyyw1e/effective-mobile-subs/pkg/logger"
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

func RegisterRoutes(r chi.Router, database *gorm.DB) {
	r.Route("/subscriptions", func(r chi.Router) {
		r.Post("/", func(w http.ResponseWriter, r *http.Request) {CreateSubHandler(w, r, database)})
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {ListSubsHandler(w, r, database)})
		r.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {GetSubHandler(w, r, database)})
		r.Put("/{id}", func(w http.ResponseWriter, r *http.Request) {UpdateSubHandler(w, r, database)})
		r.Patch("/{id}", func(w http.ResponseWriter, r *http.Request) {PatchSubHandler(w, r, database)})
		r.Delete("/{id}", func(w http.ResponseWriter, r *http.Request) {DeleteSubHandler(w, r, database)})
	})

}

func CreateSubHandler(w http.ResponseWriter, r *http.Request, database *gorm.DB) {
	req := usecase.SubscriptionReq{}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logger.Log.Errorf("failed to decode json: %v", err)
		http.Error(w, "failed to decode json", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	sub, err := usecase.CreateSub(database, req)
	if err != nil {
		logger.Log.Errorf("failed to create subscription: %v", err)
		http.Error(w, "failed to create subscription", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&sub)
}

func ListSubsHandler(w http.ResponseWriter, r *http.Request, database *gorm.DB) {
	userID := r.URL.Query().Get("user_id")
	service := r.URL.Query().Get("service")

	subs, err := usecase.ListSubs(database, userID, service)
	if err != nil {
		logger.Log.Errorf("failed to list subs: %v", err)
		http.Error(w, "failed to list subscriptions", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&subs)
}

func GetSubHandler(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	id := chi.URLParam(r, "id")
	sub, err := usecase.GetSub(db, id)
	if err != nil {
		logger.Log.Errorf("failed to get sub: %v", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
	if sub == nil {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sub)
}

func UpdateSubHandler(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	id := chi.URLParam(r, "id")
	var req usecase.SubscriptionReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	err := usecase.UpdateSub(db, id, req)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, "not found", http.StatusNotFound)
		} else {
			http.Error(w, "failed to update", http.StatusInternalServerError)
		}
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func PatchSubHandler(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	id := chi.URLParam(r, "id")
	var patch map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&patch); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	err := usecase.PatchSub(db, id, patch)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, "not found", http.StatusNotFound)
		} else {
			http.Error(w, "failed to patch", http.StatusInternalServerError)
		}
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func DeleteSubHandler(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	id := chi.URLParam(r, "id")
	err := usecase.DeleteSub(db, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, "not found", http.StatusNotFound)
		} else {
			http.Error(w, "failed to delete", http.StatusInternalServerError)
		}
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
