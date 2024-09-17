package services

import (
	"ayuphone_api/internal/models"
	"context"
	"log"
)

// CreateOrganization adds a new Organization to the database
func (t DbService) CreateOrganization(ctx context.Context, organization *models.Organization) (*models.Organization, error) {
	_, err := t.Client.DB.NewInsert().Model(organization).Returning("*").Exec(ctx)
	if err != nil {
		log.Printf("Error creating organization: %v", err)
		return nil, err
	}
	return organization, nil
}

// GetOrganization retrieves all Organizations
func (t DbService) GetOrganizations(ctx context.Context) ([]models.Organization, error) {
	var organizations []models.Organization
	err := t.Client.DB.NewSelect().Model(&organizations).Order("updated_at DESC").Scan(ctx)
	if err != nil {
		log.Printf("Error retrieving organizations %v", err)
		return nil, err
	}
	return organizations, nil
}

func (t DbService) GetOrganizationByID(ctx context.Context, id int64) (*models.Organization, error) {
	organization := new(models.Organization)
	err := t.Client.DB.NewSelect().Model(organization).Where("id = ?", id).Scan(ctx)
	if err != nil {
		log.Printf("Error retrieving organization: %v", err)
		return nil, err
	}
	return organization, nil
}

// UpdateOrganization updates an existing Organization
func (t DbService) UpdateOrganization(ctx context.Context, organization *models.Organization) error {
	_, err := t.Client.DB.NewUpdate().Model(organization).Where("id = ?", organization.ID).Exec(ctx)
	if err != nil {
		log.Printf("Error updating Organization: %v", err)
		return err
	}
	return nil
}

// DeleteOrganization deletes a Organization by ID
func (t DbService) DeleteOrganization(ctx context.Context, id int64) error {
	organization := new(models.Organization)
	_, err := t.Client.DB.NewDelete().Model(organization).Where("id = ?", id).Exec(ctx)
	if err != nil {
		log.Printf("Error deleting Organization: %v", err)
		return err
	}
	return nil
}
