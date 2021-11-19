package ports

type CustomerInterface interface {
	ValidateName()
	FormatName()
	IsDocumentCPF()
	IsDocumentCNPJ()
	ValidateDocument()
	ValidateEmail()
	GetFormatedPhone()
	ValidatePhone()
	FormatPhone()
}