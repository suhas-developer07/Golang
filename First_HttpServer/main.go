package main

import (
	"fmt"
	"net/http"  //this package in inbuilt in golang used to create an http server 
)

func main() {
	server := &http.Server{
		Addr : ":8080",
	}  //here iam establishing the http server on port 8080

  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {  //route handler 
	fmt.Fprint(w, "ok") //sending response
  })

  server.ListenAndServe()  // this line tells that listen the server in specified port 
}