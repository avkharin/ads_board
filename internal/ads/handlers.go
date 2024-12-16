package ads

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func RegisterRoutes(r *mux.Router, svc *Service) {
	r.HandleFunc("/ads", createAdHandler(svc)).Methods(http.MethodPost)
	//TODO: Register other routes
}

func createAdHandler(svc *Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var ad Ad
		if err := json.NewDecoder(r.Body).Decode(&ad); err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}

		newAd, err := svc.CreateAd(ad.Title, ad.Description, ad.Price)
		if err != nil {
			http.Error(w, "Failed to create ad", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(newAd); err != nil {
			http.Error(w, "Failed to encode ad", http.StatusInternalServerError)
		}
	}
}
