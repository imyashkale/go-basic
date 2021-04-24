package router

import (
	"log"
	"net/http"

	httprouter "github.com/julienschmidt/httprouter"
)


func FileServer(){
	router := httprouter.New()
	location := http.Dir("/home/yashkale/go/src/github.com/imyashkale/go-basic/static")
	router.ServeFiles("/static/*filepath",location)
	log.Fatal(http.ListenAndServe(":8000", router))
}