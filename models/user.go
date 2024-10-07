package models

import "encoding/json"

type User struct {
	UUIDBaseModel
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email"`
	GoogleUserId string `json:"google_user_id,omitempty"`
  TiktokUserId string `json:"tiktok_user_id,omitempty"`
}

func (u User) FullName() string {
	return u.FirstName + " " + u.LastName
}

func (u User) MarshalJSON() ([]byte, error) {
  type user User // prevent recursion
	x := user(u)
	x.GoogleUserId = ""
  x.TiktokUserId = ""
	return json.Marshal(x)
}
