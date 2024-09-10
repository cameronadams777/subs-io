package repositories

import (
	"app/database"
	"app/models"

	"github.com/google/uuid"
)

type FindUserParams struct {
	ID    string
	Email string
}

func FindUser(params FindUserParams) (*models.User, error) {
	var user models.User

  tx := database.DB

  if params.ID != "" {
    tx = tx.Where("id = ?", params.ID)
  }

  if params.Email != "" {
    tx = tx.Where("email = ?", params.Email)
  }

  err := tx.First(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func CreateUser(user models.User) error {
	err := database.DB.Create(&user).Error

	if err != nil {
		return err
	}

	return nil
}

func UpdateUser(id uuid.UUID, updated_user models.User) (*models.User, error) {
  var user_to_update models.User
  err := database.DB.First(&user_to_update, id).Error

  if err != nil {
    return nil, err
  }

  database.DB.Model(&user_to_update).Updates(updated_user)

  return &user_to_update, nil
}
