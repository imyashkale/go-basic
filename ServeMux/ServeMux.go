package servemuxex

import (
	"io"
	"log"
	"net/http"
)

func helloRoute(response http.ResponseWriter, request *http.Request) {
	io.WriteString(response, "Hello Word !! \n")

}

func StartServrMux() {

	http.HandleFunc("/", helloRoute)
	log.Fatal(http.ListenAndServe(":8000", nil))

}
