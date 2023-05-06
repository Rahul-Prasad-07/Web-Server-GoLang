package main

import (
	"fmt"
	"log"
	"net/http"
)

// Any Api or any routes have ( Req and Res)--> (w,r)

func helloHandler(w http.ResponseWriter, r *http.Request) {

	// if the path is not equal to /hello then it will return 404 not found
	if r.URL.Path != "/hello" {

		http.Error(w, "404 not found.", http.StatusNotFound)
		return

	}

	// if the method is not equal to GET then it will return Method is not supported
	if r.Method != "GET" {
		http.Error(w, "Method is not supported", http.StatusNotFound)
		return
	}

	// if the method is GET then it will return hello!
	fmt.Fprintf(w, "hello !")

}

func formHandler(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err) // if there is any error then it will return error
		return
	}
	fmt.Fprintf(w, "Post request succesfull")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)

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
