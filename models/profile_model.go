package models

type (
	Profiles []Profile

	Profile struct {
		Id      int    `json:"id"`
		Type    string `json:"type"`
		Details Detail `json:"details"`
	}

	Detail struct {
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
	}
)
