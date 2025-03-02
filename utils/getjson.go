package utils

import (
	"encoding/json"
	"io"
	"github.com/gin-gonic/gin"
)

func GetJsonBody(c *gin.Context) (map[string]interface{}, error) {
	data , e := io.ReadAll(c.Request.Body)
	var jsonData map[string]interface{}
	if e != nil {
		return jsonData, e
	}
	if e = json.Unmarshal(data, &jsonData); e != nil {
		return jsonData, e
	}
	return jsonData, nil
}