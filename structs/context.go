package structs

import "github.com/markbates/goth"

type AppContext struct {
	Key   string
	Value interface{}
}

type SessionContext struct {
	CSRFToken string
  User *goth.User
}

func (s SessionContext) IsAuthenticated() bool {
	return s.User != nil
}
