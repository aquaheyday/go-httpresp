package resp

import (
    "encoding/json"
    "net/http"
)

// Pagination represents pagination details for paginated responses.
type Pagination struct {
    Page    int `json:"page"`        // Current page number
    PerPage int `json:"per_page"`    // Number of items per page
    Total   int `json:"total"`       // Total number of items
}

// StatusResponse represents a response structure with status and message fields.
type StatusResponse struct {
    Status  string      `json:"status"`            // Status of the response (e.g., "success" or "error")
    Message string      `json:"message"`           // Message providing additional information
    Data    interface{} `json:"data,omitempty"`    // Optional data payload
}

// JSON writes a JSON response with the given status code and payload.
func JSON(w http.ResponseWriter, statusCode int, payload any) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(statusCode)
    json.NewEncoder(w).Encode(payload)
}

// Success sends a standard success response with the provided data.
func Success(w http.ResponseWriter, data any) {
    JSON(w, http.StatusOK, map[string]any{
        "success": true,
        "data":    data,
    })
}

// Error sends a standard error response with the provided message.
func Error(w http.ResponseWriter, message string) {
    JSON(w, http.StatusBadRequest, map[string]any{
        "success": false,
        "error":   message,
    })
}

// ValidationError sends a response indicating validation errors with detailed error messages.
func ValidationError(w http.ResponseWriter, errors map[string]string) {
    JSON(w, http.StatusUnprocessableEntity, map[string]any{
        "success": false,
        "errors":  errors,
    })
}

// Paginated sends a paginated response with data and pagination information.
func Paginated(w http.ResponseWriter, data any, pagination Pagination) {
    JSON(w, http.StatusOK, map[string]any{
        "success":   true,
        "data":      data,
        "pagination": pagination,
    })
}

// Custom sends a response with a custom status code and payload.
func Custom(w http.ResponseWriter, code int, payload any) {
    JSON(w, code, payload)
}

// StatusSuccess sends a response with a "success" status and the provided message and data.
func StatusSuccess(w http.ResponseWriter, message string, data interface{}) {
    response := StatusResponse{
        Status:  "success",
        Message: message,
        Data:    data,
    }
    JSON(w, http.StatusOK, response)
}

// StatusError sends a response with an "error" status and the provided message.
func StatusError(w http.ResponseWriter, message string) {
    response := StatusResponse{
        Status:  "error",
        Message: message,
    }
    JSON(w, http.StatusBadRequest, response)
}
