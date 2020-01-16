package utils

import (
	"encoding/json"
	"net/http"
)

// Message factory for custom messages
func Message(status bool, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}

// Respond object instance for headers
func Respond(w http.ResponseWriter, data map[string]interface{})  {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
