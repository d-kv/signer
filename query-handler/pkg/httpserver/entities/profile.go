package entities

type Profile struct {
	ID             string `json:"id,omitempty"`
	ProfileContent string `json:"content"`
}
