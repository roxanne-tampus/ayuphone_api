package services

import (
	"ayuphone_api/internal/models"
	"context"
	"log"
)

// GetAllRegions retrieves all regions from the database
func (s *DbService) GetAllRegions(ctx context.Context, regions *[]models.Region) error {
	err := s.Client.DB.NewSelect().Model(regions).Scan(ctx)
	if err != nil {
		log.Printf("Error retrieving regions: %v", err)
		return err
	}
	return nil
}

// GetProvincesByRegionID retrieves provinces by region ID
func (s *DbService) GetProvincesByRegionID(ctx context.Context, regionID string, provinces *[]models.Province) error {
	err := s.Client.DB.NewSelect().Model(provinces).
		Where("region_id = ?", regionID).
		Scan(ctx)
	if err != nil {
		log.Printf("Error retrieving provinces: %v", err)
		return err
	}
	return nil
}

// GetMunicipalitiesByProvinceID retrieves municipalities by province ID
func (s *DbService) GetMunicipalitiesByProvinceID(ctx context.Context, provinceID string, municipalities *[]models.Municipality) error {
	err := s.Client.DB.NewSelect().Model(municipalities).
		Where("province_id = ?", provinceID).
		Scan(ctx)
	if err != nil {
		log.Printf("Error retrieving municipalities: %v", err)
		return err
	}
	return nil
}

// GetBarangaysByMunicipalityID retrieves barangays by municipality ID
func (s *DbService) GetBarangaysByMunicipalityID(ctx context.Context, municipalityID string, barangays *[]models.Barangay) error {
	err := s.Client.DB.NewSelect().Model(barangays).
		Where("municipality_id = ?", municipalityID).
		Scan(ctx)
	if err != nil {
		log.Printf("Error retrieving barangays: %v", err)
		return err
	}
	return nil
}
