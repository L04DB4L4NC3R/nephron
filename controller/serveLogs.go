package controller

import (
	"encoding/json"
	"net/http"

	"github.com/angadsharma1016/nephron/services"
)

func (f Fetch) GetLogsJSON() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		rs := services.GetLogsJSON()
		json.NewEncoder(w).Encode(rs)
	}
}
