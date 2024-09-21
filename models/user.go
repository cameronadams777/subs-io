package models

type User struct {
	UUIDBaseModel
	FirstName           string       `json:"first_name"`
	LastName            string       `json:"last_name"`
	Email               string       `json:"email"`
}

func (u User) FullName() string {
  return u.FirstName + " " + u.LastName
}

