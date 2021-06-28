package feishu

import (
	"fmt"
	"log"
	"net/http"
)

func getBodyFromRequest(req *http.Request) []byte {
	length := req.ContentLength
	body := make([]byte, length)
	_, _ = req.Body.Read(body)
	return body
}

func httpHandler(rw http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		rw.WriteHeader(RequestHandle(getBodyFromRequest(req)))
	} else {
		rw.WriteHeader(http.StatusBadRequest)
	}
}

func Serve(host string, port int) {
	http.HandleFunc("/", httpHandler)
	err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), nil)
	if err != nil {
		log.Fatalln(err)
	}
}
