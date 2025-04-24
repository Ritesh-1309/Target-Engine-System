package handler

import (
	"encoding/json"
	"main/models"
	"main/service"
	"net/http"
)

func DeliveryHandler(campaignService service.CampaignService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		appID := query.Get("app")
		country := query.Get("country")
		os := query.Get("os")

		if appID == "" || country == "" || os == "" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "missing app param"})
			return
		}

		req := models.AppRequest{
			AppID:   appID,
			Country: country,
			OS:      os,
		}

		campaigns, err := campaignService.GetMatchingCampaigns(req)
		if err != nil {
			http.Error(w, `{"error": "internal error"}`, http.StatusInternalServerError)
			return
		}

		if len(campaigns) == 0 {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(campaigns)
	}
}
