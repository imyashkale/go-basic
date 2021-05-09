# go-basic
Routers , ServeMux and serving static files

#### Custom `ServeMux` implementation

```
// Custome Mux
type CustomeMux struct {
}

func genrateRandomNumber(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "Random Number is : %f \n", rand.Float32())
}

func (customMx *CustomeMux) ServeHTTP(response http.ResponseWriter, request *http.Request) {

	if request.URL.Path == "/" {
		genrateRandomNumber(response, request)
		return
	}
	http.NotFound(response, request)
}

func StartCustomMx() {
	mx := &CustomeMux{}
	http.ListenAndServe(":8000", mx)
}
```



#### Serving static content [ CDN ]
```
router := httprouter.New()
	location := http.Dir("/home/yashkale/go/src/github.com/imyashkale/go-basic/static")
	router.ServeFiles("/static/*filepath",location)
	log.Fatal(http.ListenAndServe(":8000", router))
 ```
