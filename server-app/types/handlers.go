package types

import (
	"encoding/json"
	"net/http"
)

type Application_t struct {
	Type   string
	Method string
}

type Response_t struct {
	StatusCode  int           `json:"status"` // Code Status
	Data        interface{}   `json:"data"`
	Message     string        `json:"message"`
	Error       Error_t       `json:"error"`
	Application Application_t `json:"application"`
}

// Marshal returns JSON-encoded response
func (r *Response_t) Marshal() ([]byte, Error_t) {
	var jsonDataByte []byte
	jsonData := map[string]interface{}{
		"status":  r.StatusCode,
		"data":    r.Data,
		"message": r.Message,
		"error": map[string]interface{}{
			"code":    r.Error.Code,
			"message": r.Error.Message,
			"result":  r.Error.Result,
		},
		"application": map[string]interface{}{
			"type": r.Application.Type,
		},
	}

	jsonDataByte, r.Error.Err = json.Marshal(jsonData)
	return jsonDataByte, r.Error
}

func (r *Response_t) Send(w http.ResponseWriter) {
	jsonDataByte, err := r.Marshal()
	err.Check(err.Err)
	w.Header().Set("Content-Type", r.Application.Type)
	w.Header().Set("Access-Control-Allow-Methods", r.Application.Method)
	w.WriteHeader(r.StatusCode)
	w.Write(jsonDataByte)
}
