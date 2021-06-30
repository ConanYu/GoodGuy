package feishu

import (
	"github.com/buger/jsonparser"
	"github.com/spf13/viper"
	"net/http"
)

func checkToken(token string, err error, body *[]byte) int {
	if err != nil || token == "" || token != viper.GetString("app.token") {
		return http.StatusBadRequest
	}
	EventHandle(body)
	return http.StatusOK
}

func requestHandleSchema1(body *[]byte) int {
	class, _ := jsonparser.GetString(*body, "type")
	if class == "url_verification" {
		return http.StatusAccepted
	}
	token, err := jsonparser.GetString(*body, "token")
	return checkToken(token, err, body)
}

func requestHandleSchema2(body *[]byte) int {
	token, err := jsonparser.GetString(*body, "header", "token")
	return checkToken(token, err, body)
}

func RequestHandle(body *[]byte) int {
	schema, err := jsonparser.GetString(*body, "schema")
	if err != nil {
		return requestHandleSchema1(body)
	}
	if schema == "2.0" {
		return requestHandleSchema2(body)
	} else {
		return http.StatusBadRequest
	}
}
