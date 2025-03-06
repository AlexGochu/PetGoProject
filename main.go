package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		log.Println("Hello World")
		fmt.Fprintln(writer, "Hello World")
		d, err := ioutil.ReadAll(request.Body)
		if err != nil {
			http.Error(writer, "Ooopsª"+err.Error(), http.StatusBadRequest)
			return
		}

		log.Println(string(d))
		fmt.Fprintln(writer, string(d))
	})

	http.HandleFunc("/goodbye", func(writer http.ResponseWriter, request *http.Request) {
		log.Println("Goodbye World")
		fmt.Fprintln(writer, "Goodbye World")
	})

	http.ListenAndServe(":8080", nil)
}
