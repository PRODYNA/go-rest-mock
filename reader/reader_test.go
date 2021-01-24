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

func Test_ReadTemplate(t *testing.T) {
	def := ReadDefinition("../test/data/template/backend.json")
	assert.NotNil(t, def)
	assert.True(t, def.Validate)
	assert.Equal(t, "test/data/template/responseTemplate.tmpl", def.Paths[0].Response.TemplateRef)
}


func Test_ReadDefinitionNoJson(t *testing.T) {
	assert.Nil(t, ReadDefinition("../test/data/error/malicious.json"))
}

func Test_ReadDefinitionFail(t *testing.T) {
	assert.Nil(t, ReadDefinition("."))
}
