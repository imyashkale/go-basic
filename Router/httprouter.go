package router

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os/exec"

	httprouter "github.com/julienschmidt/httprouter"
)

//helper function => it will not be used in route
func getComandOutput(command string, arguments ...string) string{
	cmd := exec.Command(command, arguments...)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Start()
	if err != nil {
		log.Fatal(fmt.Sprint(err) + ": " + stderr.String())
	}
	err = cmd.Wait()
	if err != nil {
		log.Fatal(fmt.Sprint(err) + ": " + stderr.String())
	}

	return out.String()
	
}

func getFileContent(response http.ResponseWriter, request *http.Request,params httprouter.Params){

	fmt.Fprintf(response, getComandOutput("/bin/cat",params.ByName("name")))
}

func goVersion(response http.ResponseWriter, request *http.Request ,params httprouter.Params){
	fmt.Fprintf(response, getComandOutput("/usr/local/go/bin/go", "version"))

}

func listFiles(response http.ResponseWriter, request *http.Request ,params httprouter.Params){
	fmt.Fprintf(response,getComandOutput("ls","-l"))
}


func HTTPRouter(){

	router := httprouter.New()

	router.GET("/api/v1/go-version",goVersion)
	router.GET("/api/v1/show-file/:name",getFileContent)
	router.GET("/api/v1/ls",listFiles)

	log.Fatal(http.ListenAndServe(":8000", router))
}