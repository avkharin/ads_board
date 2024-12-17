package ads

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func RegisterRoutes(r *mux.Router, svc *Service) {
	r.HandleFunc("/ad", createAdHandler(svc)).Methods(http.MethodPost)
	r.HandleFunc("/ads", getAllAdsHandler(svc)).Methods(http.MethodGet)
	//TODO: Register other routes
}
func getAllAdsHandler(svc *Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ads, err := svc.GetAllAds()
		if err != nil {
			http.Error(w, "Failed reading ads", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(ads); err != nil {
			http.Error(w, "Failed to encode ads", http.StatusInternalServerError)
		}
	}
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
