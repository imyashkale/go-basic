package servemuxex

import (
	"fmt"
	"log"
	"net/http"
	"math/rand"
)

func randFloat(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(response,rand.Float32())
}

func randInt(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(response, rand.Int())
}

func MultipleHandlers() {

	router := http.NewServeMux()
	router.HandleFunc("/float", randFloat)
	router.HandleFunc("/int", randInt)

	log.Fatal(http.ListenAndServe(":8000", router))
}
