package entities

type Certificate struct {
	ID                 string `json:"id,omitempty"`
	CertificateContent string `json:"content"`
}
