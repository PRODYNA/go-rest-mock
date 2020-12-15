package reader

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ReadFiles(t *testing.T) {
	assert.NotNil(t, ReadFiles("../test/data/"))
}

func Test_ReadFilesFail(t *testing.T) {
	assert.Nil(t, ReadFiles(""))
}

func Test_ReadDefinition(t *testing.T) {
	assert.NotNil(t, ReadDefinition("../test/data/backend.json"))
}

func Test_ReadDefinitionFail(t *testing.T) {
	assert.Nil(t, ReadDefinition("."))
}
