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
	def := ReadDefinition("../test/data/backend.json")
	assert.NotNil(t, def)
	assert.True(t, def.Validate)
}

func Test_ReadDefinitionFail(t *testing.T) {
	assert.Nil(t, ReadDefinition("."))
}
