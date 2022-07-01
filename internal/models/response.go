package models

import (
	"better-auth/utils"
	"time"
)

type Response struct {
	Status    string      `json:"status,omitempty,required" bson:"status,omitempty,required"`
	Message   string      `json:"message,omitempty,required" bson:"message,omitempty,required"`
	Data      interface{} `json:"data,omitempty,required" bson:"data,omitempty,required"`
	Timestamp string      `json:"timestamp,omitempty,required" bson:"timestamp,omitempty,required"`
}

func (r *Response) IsEmpty() bool {
	return r.Status == "" || r.Message == "" || r.Data == nil || r.Timestamp == ""
}

func (r *Response) SetTimestamp() {
	r.Timestamp = time.Now().Format("2006-01-02 15:04:05")
}

func CreateResponse(status string, message string, data interface{}) map[string]interface{} {
	var response Response
	response.Status = status
	response.Message = message
	response.Data = data
	response.SetTimestamp()
	return response.ToMap()
}

func (r *Response) CreateResponseString(status string, message string, data interface{}) string {
	var result string
	respmap := CreateResponse(status, message, data)
	result = utils.MapToString(respmap)
	return result
}

func (r *Response) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"status":    r.Status,
		"message":   r.Message,
		"data":      r.Data,
		"timestamp": r.Timestamp,
	}
}
