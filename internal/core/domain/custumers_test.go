package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConstant(t *testing.T) {
	assert.Equal(t, env_validate_document, "REQUIRE_DOCUMENT")
	assert.Equal(t, cpf_length, 11)
}

func TestIsDocumentCPF(t *testing.T) {
	var c = Customer{Document: 66946202848}
	x := c.IsDocumentCPF()
	assert.True(t, x)
	c = Customer{Document: 66946202818}
	x = c.IsDocumentCPF()
	assert.False(t, x)
	c = Customer{Document: 66946202840}
	x = c.IsDocumentCPF()
	assert.False(t, x)
}
