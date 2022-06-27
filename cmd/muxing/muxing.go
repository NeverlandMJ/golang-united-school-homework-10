package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

/**
Please note Start functions is a placeholder for you to start your own solution.
Feel free to drop gorilla.mux if you want and use any other solution available.

main function reads host/port from env just for an example, flavor it following your taste
*/

// Start /** Starts the web server listener on given host and port.
func Start(host string, port int) {
	router := httprouter.New()
	router.GET("/name/:name", getName)
	router.GET("/bad", badRequest)
	router.POST("/data", bodyMessage)
	router.POST("/headers", headerFunc)

	log.Println(fmt.Printf("Starting API server on %s:%d\n", host, port))
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), router); err != nil {
		log.Fatal(err)
	}
}

func headerFunc(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
	headers := r.Header
	num1, num2 := 0, 0

	if v, ok := headers[http.CanonicalHeaderKey("a")]; ok {
		num1, _ = strconv.Atoi(v[0])
	}

	if v, ok := headers[http.CanonicalHeaderKey("b")]; ok {
		num2, _ = strconv.Atoi(v[0])
	}

	res := num1+num2

	w.Header().Add("a+b", strconv.Itoa(res))
	
}

func bodyMessage(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	d, _ := io.ReadAll(r.Body)
	fmt.Fprintf(w, "I got message:\n%v", string(d)) 

}

func getName(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
	fmt.Fprintf(w, "Hello, %v!", ps.ByName("name"))
}

func badRequest(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("500 - Something bad happened!"))
}

//main /** starts program, gets HOST:PORT param and calls Start func.
func main() {
	host := os.Getenv("HOST")
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8081
	}
	Start(host, port)
}

