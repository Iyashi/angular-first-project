package main

type User struct {
	ID       uint   `json:"id"`
	PreName  string `json:"preName"`
	LastName string `json:"lastName"`
}

var users = []User{
	{1, "John", "Doe"},
	{2, "Jane", "Doe"},
	{3, "Max", "Mustermann"},
	{4, "Maxi", "Musterfrau"},
}

func FindUserByID(id uint) *User {
	for _, user := range users {
		if user.ID == id {
			return &user
		}
	}
	return nil
}
