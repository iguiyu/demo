package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, 鲑鱼~")
	})

	log.Println("Listening and serving HTTP on container port:7000")
	http.ListenAndServe(":7000", nil)
}
