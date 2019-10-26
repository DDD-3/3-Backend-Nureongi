package auth

import (
	"os"
	"context"
	"fmt"
	"strings"
	"net/http"
	jwt "github.com/dgrijalva/jwt-go"
	utils "../utils"
)

func processJwtAuthentication( next http.Handler) http.Handler {
	
	return http.HandlerFunc( func( response http.ResponseWriter, request *http.Request ) {

		noAuthList := []string{ "/api/user/register", "api/user/login" } // need make no auth url list
		requestUrl := request.URL.Path

		for _, value := range noAuthList {
			if value == requestUrl {
				next.ServeHTTP( response, request )
				return
			}
		}

		newResponse := make( map[string] interface{} )
		headerToken := request.Header.Get( "Authorization" )

		if "" == headerToken { // http 403 error
			//newResponse = utils.makeHttpMessage( false, "eErrorNotExistToken" )
			response.WriteHeader( http.StatusForbidden )
			response.Header().Add( "Content-Type", "application/json" )
			//utils.makeResponseHttpToJson( response, newResponse )
			return
		}

		splittedToken := strings.Split( headerToken, " " )

		if 2 != len( splittedToken ) { // http token length error
			//newResponse = utils.makeHttpMessage( false, "eErrorNotInvalidTokenFormat" )
			response.WriteHeader( http.StatusForbidden )
			response.Header().Add( "Content-Type", "application/json" )
			//utils.makeResponseHttpToJson( response, newResponse )
			return
		}

		
	})
}