package main

import (
	"fmt"
	"log"
	"net/http"
)

// Any Api or any routes have ( Req and Res)--> (w,r)

func helloHandler(w http.ResponseWriter, r *http.Request) {

}

func formHandler(w http.ResponseWriter, r *http.Request) {

}

func main() {

	fileServer := http.FileServer(http.Dir("./static")) // we are telling go lang that we wants to check out static directory & it knows it has to look at index.html file
	http.Handle("/", fileServer)                        // we are telling go lang that handle all the request that comes to the root dir & pass it to fileserver

	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler) // create helloHandler

	fmt.Printf("starting server at port 8080\n")

	// create server : heart of the web app : it will listen to the request & pass it to the handler
	// if err is nil then it will return nill and run server otherwise it will return error
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
