package model

type User struct {
	Xkey      string `json:"_key,omitempty"`
	ID        string `json:"id,omitempty"`
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
	Password  string `json:"password,omitempty"`
	FullName  string `json:"fullName,omitempty"`
	Time      string `json:"time,omitempty"`
}
