package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConstant(t *testing.T) {
	assert.Equal(t, env_validate_document, "REQUIRE_DOCUMENT")
	assert.Equal(t, 8,  cpf_min_length)
	assert.Equal(t, 12, cpf_max_length)
}

func TestIsValidName(t *testing.T){
	var c = Customer{}
	x := c.IsValidName()
	assert.False(t, x)
	c = Customer{Name: "Test Name"}
	x = c.IsValidName()
	assert.True(t, x)
	c = Customer{Name: ""}
	x = c.IsValidName()
	assert.False(t, x)
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
	c = Customer{Document: 1587547007}
	x = c.IsDocumentCPF()
	assert.True(t, x)
}

func TestIsDocumentCNPJ(t *testing.T) {
	var c = Customer{Document: 11222333000181}
	x := c.IsDocumentCNPJ()
	assert.True(t, x)
	c = Customer{Document: 11222333000171}
	x = c.IsDocumentCNPJ()
	assert.False(t, x)
	c = Customer{Document: 11222333000182}
	x = c.IsDocumentCNPJ()
	assert.False(t, x)
	c = Customer{Document: 74112977000137}
	x = c.IsDocumentCNPJ()
	assert.True(t, x)
}

func TestIsValidDocument (t *testing.T) {
	var c = Customer{Document: 66946202848}
	x := c.IsDocumentCPF()
	assert.True(t, x)
	c = Customer{Document: 74112977000137}
	x = c.IsDocumentCNPJ()
	assert.True(t, x)
	c = Customer{Document: 11222333000182}
	x = c.IsDocumentCNPJ()
	assert.False(t, x)
	c = Customer{Document: 66946202818}
	x = c.IsDocumentCPF()
	assert.False(t, x)
}