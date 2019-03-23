package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Nope!")
	})
	forever := make(chan struct{})
	go func() {
		checkFatal(http.ListenAndServe(":80", nil))
	}()
	go func() {
		checkFatal(http.ListenAndServeTLS(":443", "cert.pem", "key.pem", nil))
	}()
	<-forever
}

func checkFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
