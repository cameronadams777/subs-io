package models

import "github.com/google/uuid"

type Post struct {
	UUIDBaseModel
  Title string `json:"title"`
  Status string `json:"status"`
  URL string `json:"url"`
  UserID uuid.UUID `json:"user_id"`
}
