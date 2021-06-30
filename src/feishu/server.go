package feishu

import (
	"fmt"
	"github.com/buger/jsonparser"
	"log"
	"net/http"
)

func getBodyFromRequest(req *http.Request) []byte {
	length := req.ContentLength
	body := make([]byte, length)
	_, _ = req.Body.Read(body)
	return body
}

func urlVerification(rw *http.ResponseWriter, body *[]byte) {
	challenge, _ := jsonparser.GetString(*body, "challenge")
	response := fmt.Sprintf("{\"challenge\": \"%s\"}", challenge)
	_, _ = (*rw).Write([]byte(response))
}

func httpHandler(rw http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		body := getBodyFromRequest(req)
		code := RequestHandle(&body)
		if code == http.StatusAccepted {
			urlVerification(&rw, &body)
			code = http.StatusOK
		}
		rw.WriteHeader(code)
	} else {
		rw.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func Serve(host string, port int) {
	http.HandleFunc("/", httpHandler)
	err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), nil)
	if err != nil {
		log.Fatalln(err)
	}
}
