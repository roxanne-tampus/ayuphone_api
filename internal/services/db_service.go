package services

import db "ayuphone_api/internal/db"

type DbService struct {
	Client *db.Client
}

func NewDBService(client *db.Client) *DbService {
	return &DbService{Client: client}
}
