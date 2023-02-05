package dtos

type CreateUserDTO struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
}


type UpdateUserDTO struct {
	Name    string `json:"name,omitempty"`
	Surname string `json:"surname,omitempty"`
}