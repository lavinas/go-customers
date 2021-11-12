package domain

import (
	"math"
	"strconv"
	"strings"
	"net"
	"regexp"
)

const (
	env_validate_document = "REQUIRE_DOCUMENT"
	cpf_min_length        = 8
	cpf_max_length        = 12
	cnpj_min_length       = 12
	cnpj_max_length       = 16
	email_min_length      = 3
	email_max_length      = 254
	email_regex           = "^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]" +
		                    "{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"
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

// Validate customer name. 
// It should have more than a word and the first and last name should have more than a char
func (c Customer) IsValidName() bool {
	// validate if it is not blank
	if c.Name == "" {
		return false
	}
	// validade more than one name and the first and last have more then a char
	s := strings.Split(c.Name, " ")
	if len(s) < 2 {
		return false
	}
	if len(s[0]) < 2 {
		return false
	}
	if len(s[len(s)-1]) < 2 {
		return false
	}
	return true
}

// Verify if document is a valid brasilian CPF (private individual document)
// It validate two last digits with mod 11 algorithm
func (c Customer) IsDocumentCPF() bool {
	// valid is not zero
	if c.Document == 0 {
		return false
	}
	// valid length
	len := len(strconv.FormatUint(c.Document, 10))
	if len < cpf_min_length || len > cpf_max_length {
		return false
	}
	// valid check digits (2 last digits)
	dig1 := int(c.Document%100/10)
	dig2 := int(c.Document%10)
	val1 := 0
	val2 := 0
	for i := 3; i <= len; i++ {
		x := int(math.Mod(float64(c.Document), math.Pow10(i)) / math.Pow10(i-1))
		val1 += x * (i - 1)
		val2 += x * i
	}
	val2 += dig1 * 2
	val1 = int(math.Mod(float64(val1*10), float64(11)))
	val2 = int(math.Mod(float64(val2*10), float64(11)))
	return val1 == dig1 && val2 == dig2
}

// Verify if document is a valid brasilian CNPJ (legal entity document)
// It validate two last digits with mod 11 algorithm
func (c Customer) IsDocumentCNPJ() bool {
	// valid is not zero
	if c.Document == 0 {
		return false
	}
	// valid length
	len := len(strconv.FormatUint(c.Document, 10))
	if len < cnpj_min_length || len > cnpj_max_length {
		return false
	}
	// valid check digits (2 last digits)
	dig1 := int(c.Document%100/10)
	dig2 := int(c.Document%10)
	val1 := 0
	val2 := 0
	for i := 0; i <= len - 3; i++ {
		x := int(math.Mod(float64(c.Document), math.Pow10(i+3)) / math.Pow10(i+2))
		val1 += x * (int(math.Mod(float64(i), float64(8))) + 2)
		val2 += x * (int(math.Mod(float64(i+1), float64(8))) + 2)
	}
	val2 += dig1 * 2
	val1 = int(math.Mod(float64(val1), float64(11)))
	val2 = int(math.Mod(float64(val2), float64(11)))
	if val1 < 2 {
		val1 = 0
	}
	if val2 < 2{
		val2 = 0
	}
	val1 = 11 - val1
	val2 = 11 - val2
	return val1 == dig1 && val2 == dig2
}

// Validate customer document
// It should be a Brasilian CPF(private individual document) or a CNPJ(legal entity document) 
func (c Customer) IsValidDocument() bool {
	return c.IsDocumentCPF() || c.IsDocumentCNPJ()
}

// Validate customer email address string structure with go mail package
func (c Customer) IsValidEmail() bool {
		// validate if it is not blank
	if c.Email == "" {
		return false
	}
	// validate length
	e := c.Email
	if len(e) < email_min_length && len(e) > email_max_length {
		return false
	}
	// validate structure
	var emailRegex = regexp.MustCompile(email_regex)
	if !emailRegex.MatchString(e) {
		return false
	}
	parts := strings.Split(e, "@")
	mx, err := net.LookupMX(parts[1])
	if err != nil || len(mx) == 0 {
		return false
	}
	return true
}