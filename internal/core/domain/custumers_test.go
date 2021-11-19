package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func TestConstant(t *testing.T) {
	assert.Equal(t, env_validate_document, "REQUIRE_DOCUMENT")
	assert.Equal(t, 8, cpf_min_length)
	assert.Equal(t, 12, cpf_max_length)
}

func TestIsValidName(t *testing.T) {
	var c = Customer{}
	err := c.ValidateName()
	assert.NotNil(t, err)

	c = Customer{Name: "Test Name"}
	err = c.ValidateName()
	assert.Nil(t, err)
	
	c = Customer{Name: ""}
	err = c.ValidateName()
	assert.NotNil(t, err)

	c = Customer{Name: "Test"}
	err = c.ValidateName()
	assert.NotNil(t, err)
}

func TestFormatName(t *testing.T) {
	c := Customer{}
	err := c.FormatName()
	assert.NotNil(t, err)
	assert.Equal(t, "", c.Name)

	c = Customer{Name: "Test Name"}
	err = c.FormatName()
	assert.Nil(t, err)
	assert.Equal(t, "Test Name", c.Name)

	c = Customer{Name: ""}
	err = c.FormatName()
	assert.NotNil(t, err)
	assert.Equal(t, "", c.Name)

	c = Customer{Name: "test name"}
	err = c.FormatName()
	assert.Nil(t, err)
	assert.Equal(t, "Test Name", c.Name)

	c = Customer{Name: "test"}
	err = c.FormatName()
	assert.NotNil(t, err)
	assert.Equal(t, "test", c.Name)

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

func TestValidateDocument(t *testing.T) {
	var c = Customer{Document: 66946202848}
	x := c.ValidateDocument()
	assert.Nil(t, x)
	
	c = Customer{Document: 74112977000137}
	x = c.ValidateDocument()
	assert.Nil(t, x)
	
	c = Customer{Document: 11222333000182}
	x = c.ValidateDocument()
	assert.NotNil(t, x)
	
	c = Customer{Document: 66946202818}
	x = c.ValidateDocument()
	assert.NotNil(t, x)
}

func TestIsValidEmail(t *testing.T) {
	var c = Customer{Email: "teste@gmail.com"}
	x := c.ValidateEmail()
	assert.Nil(t, x)
	
	c = Customer{Email: "teste"}
	x = c.ValidateEmail()
	assert.NotNil(t, x)
	
	c = Customer{Email: "teste@"}
	x = c.ValidateEmail()
	assert.NotNil(t, x)
	
	c = Customer{Email: "teste@g"}
	x = c.ValidateEmail()
	assert.NotNil(t, x)
	
	c = Customer{Email: "teste@g.r"}
	x = c.ValidateEmail()
	assert.NotNil(t, x)
	
	c = Customer{Email: "lavinas@me.com"}
	x = c.ValidateEmail()
	assert.Nil(t, x)
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
	x := c.ValidatePhone()
	assert.Nil(t, x)

	c = Customer{PhoneNumber: 97776755}
	x = c.ValidatePhone()
	assert.NotNil(t, x)
	
	c = Customer{PhoneNumber: 551197776755}
	x = c.ValidatePhone()
	assert.Nil(t, x)
	
	c = Customer{PhoneNumber: 551197776755, PhoneCountry: "US"}
	x = c.ValidatePhone()
	assert.NotNil(t, x)
	
	c = Customer{PhoneNumber: 2015550123, PhoneCountry: "US"}
	x = c.ValidatePhone()
	assert.Nil(t, x)
	
	c = Customer{PhoneNumber: 12015550123, PhoneCountry: "US"}
	x = c.ValidatePhone()
	assert.Nil(t, x)
	
	c = Customer{PhoneNumber: 7400123456, PhoneCountry: "GB"}
	x = c.ValidatePhone()
	assert.Nil(t, x)
	
	c = Customer{PhoneNumber: 447400123456, PhoneCountry: "GB"}
	x = c.ValidatePhone()
	assert.Nil(t, x)
	
	c = Customer{PhoneNumber: 1197776755, PhoneCountry: "XY"}
	x = c.ValidatePhone()
	assert.NotNil(t, x)
}

func TestFormatPhone(t *testing.T) {
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

func TestNewCustomer(t *testing.T) {
	c := NewCustomer()
	_, err := uuid.Parse(c.Id)
	var zeroUInt64 uint64 = 0 
	assert.Nil(t, err)
	assert.Equal(t, c.Name, "")
	assert.Equal(t, c.Document, zeroUInt64)
	assert.Equal(t, c.Email, "")
	assert.Equal(t, c.PhoneCountry, "")
	assert.Equal(t, c.PhoneNumber, zeroUInt64)
	assert.Equal(t, c.Password, "")
}

func TestIsPasswordCrypted(t *testing.T) {
	var c = Customer{Password: "xxxsasssadsa"}
	i := c.IsPasswordCrypted()
	assert.False(t, i)
	c.FormatPassword()
	i = c.IsPasswordCrypted()
	assert.True(t, i)
	
}

func TestValidatePassword(t *testing.T) {
	var c = Customer{}
	err := c.ValidatePassword()
	assert.NotNil(t, err)
	c = Customer{Password: "uwquweeuwuewue"}
	err = c.ValidatePassword()
	assert.Nil(t, err)
}


func TestFormatPassword(t *testing.T) {
	var c = Customer{Password: "xxxsasssadsa"}
	err := c.FormatPassword()
	assert.Nil(t, err)
	err = bcrypt.CompareHashAndPassword([]byte(c.Password), []byte("xxxsasssadsa"))
	assert.Nil(t, err)
	err = bcrypt.CompareHashAndPassword([]byte(c.Password), []byte("xxxsasssadsab"))
	assert.NotNil(t, err)
	err = c.FormatPassword()
	assert.NotNil(t, err)
}




