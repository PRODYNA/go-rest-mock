package reader

import (
	"encoding/json"
	"github.com/prodyna/go-rest-mock/model"
	"io/ioutil"
	"log"
	"os"
)

// ReadDefinition reads a mock definition
func ReadDefinition(path string) *model.MockDefinition {

	data, err := ioutil.ReadFile(path)

	if err != nil {
		log.Fatal(err)
	}

	var mockDefinition model.MockDefinition

	err = json.Unmarshal(data, &mockDefinition)
	if err != nil {
		log.Fatal(err)
	}

	return &mockDefinition
}

// ReadFiles reads the config files in the given directory.
func ReadFiles(path string) []os.FileInfo {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	return files
}
