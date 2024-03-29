package entities

type Device struct {
	Identifier string `json:"identifier"`
	Name       string `json:"name"`
	UserId     string `json:"userId"`
}
