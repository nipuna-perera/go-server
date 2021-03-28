package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	uril := fmt.Sprintf("Request URI: %v", r.RequestURI)
	log.Println(uril)
	fmt.Fprintf(w, uril)
}


func serveFile(w http.ResponseWriter, r *http.Request){
	ldir, _:=os.Getwd()

	http.ServeFile(w, r, ldir+r.URL.Path)
}
func main() {
	http.HandleFunc("/hello", sayHello)
	http.HandleFunc("/", serveFile)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
