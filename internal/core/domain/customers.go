package domain

import (
	"errors"
	"math"
)

const (
	env_validate_document = "REQUIRE_DOCUMENT"
	cpf_length        = 11
)

// Customer represents basic informations of a Costumer
//
// A Custumer is the security principal for this application.
// It's also used as one of main axes for reporting.
//
// A Constumer can have links with whom they can be connected in some form.
//
// swagger:model
type Customer struct {
	// the id for this client
	// required: true
	Id uint64 `json:"id"`
	// the name for this client
	// required: true
	Name string `json:"name"`
	// the document number for this client
	// required: false
	Document uint64 `json:"document"`
	// the email address for this client
	// required: true
	// example: user@provider.net
	Email string `json:"email"`
	// the cell number for this client in the i164 pattern
	// required: false
	Phone uint64 `json:"phone"`
	// the unified password for this client
	// required: false
	Password string `json:"password"`
}

// Validate is a Client method that validate if the fields are in the right expected format
func (c Customer) Validate() error {
	return nil
}

// Validate Custumer Name
func (c Customer) ValidateName() error {
	if c.Name == "" {
		return errors.New("name is empty")
	}
	return nil
}

// Verify if document is a valid brasilian CPF
func (c Customer) IsDocumentCPF() bool {
	// valid is not zero
	if c.Document == 0 {
		return false
	}
	// valid digit 1
	dig1 := int(c.Document%100/10)
	dig2 := int(c.Document%10)
	val1 := 0
	val2 := 0
	for i := 3; i <= cpf_length; i++ {
		x := int(math.Mod(float64(c.Document), math.Pow10(i)) / math.Pow10(i-1))
		val1 += x * (i - 1)
		val2 += x * i
	}
	val2 += dig1 * 2
	val1 = int(math.Mod(float64(val1*10), float64(11)))
	val2 = int(math.Mod(float64(val2*10), float64(11)))
	return val1 == dig1 && val2 == dig2
}
