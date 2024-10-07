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

	// Group devices by brand and include their IDs and Models
	brandModelsMap := make(map[string][]models.DeviceModel)
	for _, device := range devices {
		brandModelsMap[device.Brand] = append(brandModelsMap[device.Brand], models.DeviceModel{
			ID:    device.ID,
			Model: device.Model,
		})
	}

	// Convert the map to a slice of models.BrandModel
	var result []models.BrandModel
	for brand, deviceModels := range brandModelsMap {
		result = append(result, models.BrandModel{
			Brand:  brand,
			Models: deviceModels,
		})
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
