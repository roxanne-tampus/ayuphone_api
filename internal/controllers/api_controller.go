package controllers

import (
	db_client "ayuphone_api/internal/db"
	"ayuphone_api/internal/services"
)

type ApiController struct {
	DbClient  *db_client.Client
	DbService *services.DbService
}
