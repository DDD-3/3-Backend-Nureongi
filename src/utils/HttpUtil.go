package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func makeHttpMessage(status bool, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}

func makeResponseHttpToJson(response http.ResponseWriter, statusCode int, data map[string]interface{}) {

	response.Header().Add("Content-Type", "application/json")
	response.WriteHeader(statusCode)

	err := json.NewEncoder(response).Encode(data)
	if err != nil {
		fmt.Fprintf(response, "%s", err.Error())
	}
}

func makeErrorResponseHttp(response http.ResponseWriter, statusCode int, err error) {

}
