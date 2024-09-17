package services

import (
	"ayuphone_api/internal/models"
	"context"
)

func (dbs *DbService) GetAllDevice(ctx context.Context) ([]models.BrandModel, error) {
	var devices []models.Device
	err := dbs.Client.DB.NewSelect().
		Model(&devices).
		Scan(ctx)
	if err != nil {
		return nil, err
	}

	// Group devices by brand
	brandModelsMap := make(map[string][]string)
	for _, device := range devices {
		brandModelsMap[device.Brand] = append(brandModelsMap[device.Brand], device.Model)
	}

	// Convert the map to a slice of models.BrandModel
	var result []models.BrandModel
	var res models.BrandModel
	for brand, models := range brandModelsMap {
		res.Brand = brand
		res.Models = models
		result = append(result, res)
	}
	return result, nil
}

func (dbs *DbService) GetAllDeviceIssues(ctx context.Context) ([]models.DeviceIssue, error) {
	var issues []models.DeviceIssue
	err := dbs.Client.DB.NewSelect().Model(&issues).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return issues, nil
}
