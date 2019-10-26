package utils


import (
	"net/http"
	"encoding/json"
)

func makeHttpMessage( status bool, message string ) ( map[string]interface{} ) {
	return map[string]interface{} { "status": status, "message": message }
}

func makeResponseHttpToJson( response http.ResponseWriter, data map[string]interface{} ) {
	response.Header().Add( "Content-Type", "application/json" )
	json.NewEncoder( response ).Encode( data )
}