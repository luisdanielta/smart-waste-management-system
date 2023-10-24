package handlers

import (
	"encoding/json"
)

type Error_t struct {
	Code    int // Code error log
	Message string
	Result  bool
	Err     error
}

type Application_t struct {
	Type   string
	Url    string
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
		"status": r.StatusCode,
		"data":   r.Data,
		"error": map[string]interface{}{
			"code":    r.Error.Code,
			"message": r.Error.Message,
			"result":  r.Error.Result,
		},
		"application": map[string]interface{}{
			"type":   r.Application.Type,
			"url":    r.Application.Url,
			"method": r.Application.Method,
		},
	}

	jsonDataByte, r.Error.Err = json.Marshal(jsonData)
	return jsonDataByte, r.Error
}
