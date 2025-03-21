package services

import (
	"ayuphone_api/internal/models"
	"context"
	"fmt"
)

func (u DbService) CreateUser(ctx context.Context, user *models.User) (int64, error) {
	_, err := u.Client.DB.NewInsert().Model(user).Returning("id").Exec(ctx)
	if err != nil {
		return 0, fmt.Errorf("failed to create user: %w", err)
	}
	return user.ID, nil
}

func (u DbService) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	user := &models.User{}
	err := u.Client.DB.NewSelect().Model(user).Where("email = ?", email).Scan(context.Background())
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return user, nil
}

func (u DbService) GetUserByPhoneNumber(ctx context.Context, phoneNumber string) (*models.User, error) {
	user := new(models.User)
	err := u.Client.DB.NewSelect().Model(user).Where("phone_number = ?", phoneNumber).Scan(ctx)
	return user, err
}

func (u DbService) GetUserByID(ctx context.Context, userID int64) (*models.User, error) {
	var user models.User
	err := u.Client.DB.NewSelect().Model(&user).Where("id = ?", userID).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u DbService) GetUsers(ctx context.Context) (*[]models.User, error) {
	user := &[]models.User{}
	err := u.Client.DB.NewSelect().Model(user).Scan(context.Background())
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return user, nil
}
