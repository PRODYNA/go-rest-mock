package main

import (
	"github.com/prodyna/go-rest-mock/handler"
	"github.com/prodyna/go-rest-mock/model"
	"github.com/prodyna/go-rest-mock/reader"
	"log"
	"net/http"
)

func main() {

	files := reader.ReadFiles("./test/data")
	size := len(files)
	if size == 0 {
		log.Println("No mock definitions found")
		return
	}
	for i, file := range files {

		md := reader.ReadDefinition("./test/data/" + file.Name())

		if i == size-1 {
			// last one blocks
			runServer(md)
		} else {
			// using non blocking listen & serve
			go func() {
				runServer(md)
			}()
		}
	}
}

func runServer(md *model.MockDefinition) {
	log.Println("Starting mock on port:", md.Port, "for backend:", md.ID)
	log.Fatal(http.ListenAndServe(":"+md.Port, handler.NewHandler(md)))
}
