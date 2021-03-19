package main

import (
	"github.com/placons/go-rest-mock/config"
	"github.com/placons/go-rest-mock/handler"
	"github.com/placons/go-rest-mock/model"
	"github.com/placons/go-rest-mock/reader"
	"log"
	"net/http"
	"path/filepath"
)

func main() {

	cfg := config.Parse()

	files := reader.ReadFiles(cfg.Path)
	size := len(files)
	if size == 0 {
		fullPath, err := filepath.Abs(cfg.Path)
		if err == nil {
			log.Println("No mock definitions found in path " + fullPath)
		} else {
			log.Println("No mock definitions found in path " + cfg.Path)
		}
		return
	}
	for i, file := range files {

		if file.IsDir() {
			continue
		}

		md := reader.ReadDefinition(cfg.Path + "/" + file.Name())

		if i == size-1 {
			// last one blocks and prevents from exiting
			runServer(md, cfg)
		} else {
			// using non blocking listen & serve
			go func() {
				runServer(md, cfg)
			}()
		}
	}
}

func runServer(md *model.MockDefinition, cfg *config.Config) {
	log.Println("Starting mock on port:", md.Port, "for backend:", md.ID)
	log.SetFlags(log.Llongfile)
	if md.TLS {
		log.Fatal(http.ListenAndServeTLS(":"+md.Port, "server.crt", "server.key", handler.NewHandler(md, cfg)))
	} else {
		log.Fatal(http.ListenAndServe(":"+md.Port, handler.NewHandler(md, cfg)))
	}
}
