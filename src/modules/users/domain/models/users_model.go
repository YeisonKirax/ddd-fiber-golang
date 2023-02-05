package models

type User struct {
	Id      string `json:"id,omitempty" bson:"_id,omitempty"`
	Name    string `json:"name,omitempty"`
	Surname string `json:"surname,omitempty"`
}

// Newtype returns new type.
func NewUser(name string, surname string) User {

	return User{Name: name, Surname: surname}
}
