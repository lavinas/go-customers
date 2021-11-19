package domain

import (
	"errors"
	"math"
	"net"
	"regexp"
	"strconv"
	"strings"

	"github.com/dongri/phonenumber"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
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
	default_country = "BR"
)

// Customer represents basic informations of a Costumer
//
// A Custumer is the security principal for this application.
// It's also used as one of main axes for reporting.
//
// A Customer can have links with whom they can be connected in some form.
//
// swagger:model
type Customer struct {
	// the id for this customer
	// required: true
	Id string `json:"id"`
	// the name for this customer
	// required: true
	Name string `json:"name"`
	// the document number for this customer
	// required: depends on system configuration
	Document uint64 `json:"document"`
	// the email address for this custumer
	// required: depends on system configuration
	// example: user@provider.net
	Email string `json:"email"`
	// the phone country of this customer(uses ISO3166 country code with two digits)
	// required: false
	// it's assume default country of running system
	PhoneCountry string `json:"phone_country"`
	// the phone number for this customer
	// required: depends on system configuration
	PhoneNumber uint64 `json:"phone_number"`
	// the unified master password for this customer
	// required: depends on system configuration
	Password string `json:"password"`
}

func NewCustomer() *Customer {
	id := uuid.New()
	return &Customer{Id: id.String()}
}

// Validate customer name.
// It should have more than a word and the first and last name should have more than a char
func (c *Customer) ValidateName() error {
	// validate if it is not blank
	if c.Name == "" {
		return errors.New("name should no be Nil")
	}
	// validade more than one name and the first and last have more then a char
	s := strings.Split(c.Name, " ")
	if len(s) < 2 {
		return errors.New("name should have at least two words")
	}
	if len(s[0]) < 2 {
		return errors.New("nirst name should have at least two letters")
	}
	if len(s[len(s)-1]) < 2 {
		return errors.New("last name should have at least two letters")
	}
	return nil
}

// Format Name of custumer characteres
func (c *Customer) FormatName() error {
	if err := c.ValidateName(); err != nil {
		return err
	}
	c.Name = strings.Title(strings.ToLower(c.Name))
	return nil
} 

// Verify if document is a valid brasilian CPF (private individual document)
// It validate two last digits with mod 11 algorithm
func (c *Customer) IsDocumentCPF() bool {
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
	dig1 := int(c.Document % 100 / 10)
	dig2 := int(c.Document % 10)
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
	if val1 != dig1 || val2 != dig2 {
		return  false
	}
	// Ok
	return true
}

// Verify if document is a valid brasilian CNPJ (legal entity document)
// It validate two last digits with mod 11 algorithm
func (c *Customer) IsDocumentCNPJ() bool {
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
	dig1 := int(c.Document % 100 / 10)
	dig2 := int(c.Document % 10)
	val1 := 0
	val2 := 0
	for i := 0; i <= len-3; i++ {
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
	if val2 < 2 {
		val2 = 0
	}
	val1 = 11 - val1
	val2 = 11 - val2
	if val1 != dig1 || val2 != dig2 {
		return false
	}
	// Ok
	return true
}

// Validate customer document
// It should be a Brasilian CPF(private individual document) or a CNPJ(legal entity document)
func (c *Customer) ValidateDocument() error {
	if !c.IsDocumentCPF() && !c.IsDocumentCNPJ() {
		return errors.New("document should have a CPF or CNPJ number")
	}
	return nil
}

// Validate customer email address string
func (c *Customer) ValidateEmail() error {
	// validate if it is not blank
	if c.Email == "" {
		return errors.New("email should have a valid email address format")
	}
	// validate length
	e := c.Email
	if len(e) < email_min_length && len(e) > email_max_length {
		return errors.New("email should have a valid email address format")
	}
	// validate structure
	var emailRegex = regexp.MustCompile(email_regex)
	if !emailRegex.MatchString(e) {
		return errors.New("email should have a valid email address format")
	}
	parts := strings.Split(e, "@")
	mx, err := net.LookupMX(parts[1])
	if err != nil || len(mx) == 0 {
		return errors.New("email should have a valid email address format")
	}
	return nil
}

// Return full valid number and country code of the number
// if customer has a valid number and/or country
// If has an invalid number and/or contry than (0, "") is returned
func (c *Customer) GetFormatedPhone() (uint64, string) {
	// check non nil number
	if c.PhoneNumber == 0 {
		return 0, ""
	}
	// try to find country of phone number
	country := c.PhoneCountry
	if country == "" {
		iso := phonenumber.GetISO3166ByNumber(strconv.FormatUint(c.PhoneNumber, 10), false)
		if iso.Alpha2 == "" {
			country = default_country
		} else {
			country = iso.Alpha2
		}
	}
	// try to format number
	var u uint64
	var err error
	var sn string = phonenumber.Parse(strconv.FormatUint(c.PhoneNumber, 10), country)
	u, err = strconv.ParseUint(sn, 10, 64)
	if err != nil {
		return 0, ""
	}
	return u, country
}

// Validate phone number or/and country
func (c *Customer) ValidatePhone() error {
	var nilNum uint64 = 0
	n, _ := c.GetFormatedPhone()
	if n == nilNum {
		return errors.New("PhoneNumber should have a valid phone number format")
	}
	return nil
}

// Replace number and country phone of custumer with country prefix if it has a valid number
// if there is an error in phone number or country, both will be replaced to nil
func (c *Customer) FormatPhone() {
	if c.PhoneNumber == 0 {
		c.PhoneCountry = ""
		return
	}
	number, country := c.GetFormatedPhone()
	if number == 0 {
		c.PhoneCountry = ""
		c.PhoneNumber = 0
		return
	}
	c.PhoneNumber = number
	c.PhoneCountry = country
}

func (c *Customer) IsPasswordCrypted() bool {
	if c.Password == "" {
		return false
	}
	cost, err := bcrypt.Cost([]byte(c.Password))
	if err != nil {
		return false
	}
	return cost == bcrypt.DefaultCost
}

func (c *Customer) ValidatePassword() error {
	if c.Password == "" {
		return errors.New("password should not by empty")
	}
	if c.IsPasswordCrypted() {
		return errors.New("password already crypted")
	}
	return nil
}

func (c *Customer) FormatPassword() error {
	if err := c.ValidatePassword(); err != nil {
		return err
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(c.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	c.Password = string(hash)
	return nil
}
