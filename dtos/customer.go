package dtos

type CustomerDTO struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number,omitempty"`
	Address     string `json:"address,omitempty"`
	Password    string `json:"password,omitempty"`
}
