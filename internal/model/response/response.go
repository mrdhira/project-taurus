package response

import (
	"encoding/json"
	"net/http"
)

type BaseResponse struct {
	Message string      `json:"message"`
	Error   any         `json:"error"`
	Data    interface{} `json:"data"`
}

func NewHttpJSONResponse(w http.ResponseWriter, statusCode int, responseMessage string, err any, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	resp := BaseResponse{
		Message: responseMessage,
		Error:   err,
		Data:    data,
	}
	json.NewEncoder(w).Encode(resp)
}
