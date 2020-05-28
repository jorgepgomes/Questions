package model

import "encoding/json"

func ToJson(data interface{}) string {
	return ToJsonString(data)
}
func ToJsonString(data interface{}) string {
	return string(ToJsonBytes(data))
}
func ToJsonBytes(data interface{}) []byte {
	b, err := json.Marshal(data)
	if err != nil {
		return []byte("")
	} else {
		return b
	}
}

type Response struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}
