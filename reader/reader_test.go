package reader

import "testing"

func Test_ReadFiles(t *testing.T) {
	ReadFiles("../test/data/")
}


func Test_ReadDefinition(t *testing.T) {
	ReadDefinition("../test/data/backend.json")
}
