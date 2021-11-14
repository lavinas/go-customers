package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"	
)

func TestConstant(t *testing.T) {
	assert.Equal(t, env_validate_document, "REQUIRE_DOCUMENT")
	assert.Equal(t, 8, cpf_min_length)
	assert.Equal(t, 12, cpf_max_length)
}

func TestIsValidName(t *testing.T) {
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

func TestIsValidDocument(t *testing.T) {
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

func TestIsValidEmail(t *testing.T) {
	var c = Customer{Email: "teste@gmail.com"}
	x := c.IsValidEmail()
	assert.True(t, x)
	
	c = Customer{Email: "teste"}
	x = c.IsValidEmail()
	assert.False(t, x)
	
	c = Customer{Email: "teste@"}
	x = c.IsValidEmail()
	assert.False(t, x)
	
	c = Customer{Email: "teste@g"}
	x = c.IsValidEmail()
	assert.False(t, x)
	
	c = Customer{Email: "teste@g.r"}
	x = c.IsValidEmail()
	assert.False(t, x)
	
	c = Customer{Email: "lavinas@me.com"}
	x = c.IsValidEmail()
	assert.True(t, x)
}

func TestGetFormatedPhone(t *testing.T) {
	var c Customer
	var number uint64
	var country string
	var expNumber uint64
	var expCountry string

	c = Customer{PhoneNumber: 1197776755}
	number, country = c.GetFormatedPhone()
	expNumber = 551197776755
	expCountry = "BR"
	assert.Equal(t, expNumber, number)
	assert.Equal(t, expCountry, country)

	c = Customer{PhoneNumber: 97776755}
	number, country = c.GetFormatedPhone()
	expNumber = 0
	expCountry = ""
	assert.Equal(t, expNumber, number)
	assert.Equal(t, expCountry, country)

	c = Customer{PhoneNumber: 551197776755}
	number, country = c.GetFormatedPhone()
	expNumber = 551197776755
	expCountry = "BR"
	assert.Equal(t, expNumber, number)
	assert.Equal(t, expCountry, country)

	c = Customer{PhoneNumber: 551197776755, PhoneCountry: "US"}
	number, country = c.GetFormatedPhone()
	expNumber = 0
	expCountry = ""
	assert.Equal(t, expNumber, number)
	assert.Equal(t, expCountry, country)

	c = Customer{PhoneNumber: 2015550123, PhoneCountry: "US"}
	number, country = c.GetFormatedPhone()
	expNumber = 12015550123
	expCountry = "US"
	assert.Equal(t, expNumber, number)
	assert.Equal(t, expCountry, country)

	c = Customer{PhoneNumber: 12015550123, PhoneCountry: "US"}
	number, country = c.GetFormatedPhone()
	expNumber = 12015550123
	expCountry = "US"
	assert.Equal(t, expNumber, number)
	assert.Equal(t, expCountry, country)

	c = Customer{PhoneNumber: 7400123456, PhoneCountry: "GB"}
	number, country = c.GetFormatedPhone()
	expNumber = 447400123456
	expCountry = "GB"
	assert.Equal(t, expNumber, number)
	assert.Equal(t, expCountry, country)

	c = Customer{PhoneNumber: 447400123456, PhoneCountry: "GB"}
	number, country = c.GetFormatedPhone()
	expNumber = 447400123456
	expCountry = "GB"
	assert.Equal(t, expNumber, number)
	assert.Equal(t, expCountry, country)

	c = Customer{PhoneNumber: 1197776755, PhoneCountry: "XY"}
	number, country = c.GetFormatedPhone()
	expNumber = 0
	expCountry = ""
	assert.Equal(t, expNumber, number)
	assert.Equal(t, expCountry, country)

	c = Customer{PhoneNumber: 551197776755}
	number, country = c.GetFormatedPhone()
	expNumber = 551197776755
	expCountry = "BR"
	assert.Equal(t, expNumber, number)
	assert.Equal(t, expCountry, country)

	c = Customer{PhoneNumber: 447400123456}
	number, country = c.GetFormatedPhone()
	expNumber = 447400123456
	expCountry = "GB"
	assert.Equal(t, expNumber, number)
	assert.Equal(t, expCountry, country)
}

func TestIsValidaPhone(t *testing.T) {
	var c = Customer{PhoneNumber: 1197776755}
	x := c.IsValidPhone()
	assert.True(t, x)

	c = Customer{PhoneNumber: 97776755}
	x = c.IsValidPhone()
	assert.False(t, x)
	
	c = Customer{PhoneNumber: 551197776755}
	x = c.IsValidPhone()
	assert.True(t, x)
	
	c = Customer{PhoneNumber: 551197776755, PhoneCountry: "US"}
	x = c.IsValidPhone()
	assert.False(t, x)
	
	c = Customer{PhoneNumber: 2015550123, PhoneCountry: "US"}
	x = c.IsValidPhone()
	assert.True(t, x)
	
	c = Customer{PhoneNumber: 12015550123, PhoneCountry: "US"}
	x = c.IsValidPhone()
	assert.True(t, x)
	
	c = Customer{PhoneNumber: 7400123456, PhoneCountry: "GB"}
	x = c.IsValidPhone()
	assert.True(t, x)
	
	c = Customer{PhoneNumber: 447400123456, PhoneCountry: "GB"}
	x = c.IsValidPhone()
	assert.True(t, x)
	
	c = Customer{PhoneNumber: 1197776755, PhoneCountry: "XY"}
	x = c.IsValidPhone()
	assert.False(t, x)
}

func TestFormat(t *testing.T) {
	var c = Customer{PhoneNumber: 1197776755}
	var u uint64
	c.FormatPhone()
	u = 551197776755
	assert.Equal(t, "BR", c.PhoneCountry)
	assert.Equal(t, u, c.PhoneNumber)
	c = Customer{PhoneNumber: 97776755}
	c.FormatPhone()
	u = 0
	assert.Equal(t, "", c.PhoneCountry)
	assert.Equal(t, u, c.PhoneNumber)
	c = Customer{PhoneNumber: 551197776755}
	c.FormatPhone()
	u = 551197776755
	assert.Equal(t, "BR", c.PhoneCountry)
	assert.Equal(t, u, c.PhoneNumber)
	c = Customer{PhoneNumber: 551197776755, PhoneCountry: "US"}
	c.FormatPhone()
	u = 0
	assert.Equal(t, "", c.PhoneCountry)
	assert.Equal(t, u, c.PhoneNumber)
	c = Customer{PhoneNumber: 2015550123, PhoneCountry: "US"}
	c.FormatPhone()
	u = 12015550123
	assert.Equal(t, "US", c.PhoneCountry)
	assert.Equal(t, u, c.PhoneNumber)
	c = Customer{PhoneNumber: 12015550123, PhoneCountry: "US"}
	c.FormatPhone()
	u = 12015550123
	assert.Equal(t, "US", c.PhoneCountry)
	assert.Equal(t, u, c.PhoneNumber)
	c = Customer{PhoneNumber: 7400123456, PhoneCountry: "GB"}
	c.FormatPhone()
	u = 447400123456
	assert.Equal(t, "GB", c.PhoneCountry)
	assert.Equal(t, u, c.PhoneNumber)
	c = Customer{PhoneNumber: 447400123456, PhoneCountry: "GB"}
	c.FormatPhone()
	u = 447400123456
	assert.Equal(t, "GB", c.PhoneCountry)
	assert.Equal(t, u, c.PhoneNumber)
	c = Customer{PhoneNumber: 1197776755, PhoneCountry: "XY"}
	c.FormatPhone()
	u = 0
	assert.Equal(t, "", c.PhoneCountry)
	assert.Equal(t, u, c.PhoneNumber)
}

