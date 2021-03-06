package utils

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Message(status bool, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}

func Respond(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func ErrorBadRequest(w http.ResponseWriter, message string) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(map[string]interface{}{"error": message})
}
func GErrorBadRequest(c *gin.Context, message string) {
	c.JSON(http.StatusBadRequest, gin.H{"error": message})
}

func ErrorNotFound(w http.ResponseWriter, message string) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]interface{}{"error": message})
}
func GErrorNotFound(c *gin.Context, message string) {
	c.JSON(http.StatusNotFound, gin.H{"error": message})
}
func ErrorInternal(w http.ResponseWriter, message string) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(map[string]interface{}{"error": message})
}
func GErrorInternal(c *gin.Context, message string) {
	c.JSON(http.StatusInternalServerError, gin.H{"error": message})
}

func Success(w http.ResponseWriter, data interface{}) {
	Respond(w, map[string]interface{}{"result": data})
}

func GSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, data)
}
