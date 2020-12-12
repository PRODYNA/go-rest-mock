package main

import (
	"github.com/prodyna/go-rest-mock/config"
	"github.com/prodyna/go-rest-mock/handler"
	"github.com/prodyna/go-rest-mock/model"
	"github.com/prodyna/go-rest-mock/reader"
	"log"
	"net/http"
)

func main() {

	config := config.Parse()

	files := reader.ReadFiles(config.Path)
	size := len(files)
	if size == 0 {
		log.Println("No mock definitions found in path " + config.Path)
		return
	}
	for i, file := range files {

		md := reader.ReadDefinition(config.Path + "/" + file.Name())

		if i == size-1 {
			// last one blocks and prevents from exiting
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
