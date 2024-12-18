package ads

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func RegisterRoutes(r *mux.Router, svc *Service) {
	r.HandleFunc("/ad", createAdHandler(svc)).Methods(http.MethodPost)
	r.HandleFunc("/ad", getAdHandler(svc)).Methods(http.MethodGet)
	r.HandleFunc("/ads", getAllAdsHandler(svc)).Methods(http.MethodGet)
	r.HandleFunc("/ad", updateAdHandler(svc)).Methods(http.MethodPut)
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

func getAdHandler(svc *Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.URL.Query().Get("id")
		if idStr == "" {
			http.Error(w, "missing 'id' parameter", http.StatusBadRequest)
			return
		}

		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid 'id' parameter", http.StatusBadRequest)
			return
		}

		ad, err := svc.GetAd(id)
		if err != nil {
			http.Error(w, "Ad not found", http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(*ad); err != nil {
			http.Error(w, "Failed to encode ad", http.StatusInternalServerError)
		}
	}
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

func updateAdHandler(svc *Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var ad Ad
		if err := json.NewDecoder(r.Body).Decode(&ad); err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}

		ok, err := svc.UpdateAd(ad)
		if err != nil || !ok {
			http.Error(w, "Failed update Ad", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
