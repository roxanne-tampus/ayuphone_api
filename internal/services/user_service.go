package services

import (
	db "ayuphone_api/internal/db"
	"ayuphone_api/internal/models"
	"context"
	"fmt"
)

func CreateUser(ctx context.Context, user *models.User) error {
	_, err := db.DBClient.NewInsert().Model(user).Exec(ctx)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	return nil
}

func GetUserByEmail(email string) (*models.User, error) {
	user := &models.User{}
	err := db.DBClient.NewSelect().Model(user).Where("email = ?", email).Scan(context.Background())
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return user, nil
}

func GetUserByID(ctx context.Context, userID int64) (*models.User, error) {
	var user models.User
	err := db.DBClient.NewSelect().Model(&user).Where("id = ?", userID).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
