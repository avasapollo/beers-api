package utils

import (
	"encoding/json"
	"net/http"
)

func ParseJsonBodyRequest(r *http.Request, target interface{}) error {
	var body []byte
	_, err := r.Body.Read(body)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	if err := json.Unmarshal(body, target); err != nil {
		return err
	}

	return nil
}

func RespondWithError(w http.ResponseWriter, code int, message string) {
	RespondWithJSON(w, code, map[string]string{"error": message})
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
