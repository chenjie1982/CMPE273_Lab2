package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"encoding/json"
)
type Greet struct {
	Greeting string `json:"greeting"`
}

type Name struct {
	Name string `json:"name"`
}

func hello(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
	fmt.Fprintf(rw, "Hello, %s!\n", p.ByName("name"))
}

func helloGreet(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
	var args Name

	err := json.NewDecoder(req.Body).Decode(&args)
	if err != nil {
		fmt.Println("Error! Can't unmarshal Json from create emulator request.", req.Body)
		return
	}
	var res Greet
	res.Greeting = "Hello,"+args.Name+"!"

	if err := json.NewEncoder(rw).Encode(res); err != nil {
		panic(err)
	}
}

func main() {
	mux := httprouter.New()
	mux.GET("/hello/:name", hello)
	mux.POST("/hello/", helloGreet)
	server := http.Server{
	Addr:        "0.0.0.0:8080",
	Handler: mux,
	}
	server.ListenAndServe()
}