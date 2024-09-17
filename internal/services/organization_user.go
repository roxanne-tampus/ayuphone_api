package services

import (
	"ayuphone_api/internal/models"
	"context"
	"log"
)

// CreateOrganizationUser adds a new OrganizationUser to the database
func (t DbService) CreateOrganizationUser(ctx context.Context, organizationUser *models.OrganizationUser) (*models.OrganizationUser, error) {
	_, err := t.Client.DB.NewInsert().Model(organizationUser).Returning("*").Exec(ctx)
	if err != nil {
		log.Printf("Error creating organizationUser: %v", err)
		return nil, err
	}
	return organizationUser, nil
}

// GetOrganizationUser retrieves all OrganizationUsers
func (t DbService) GetOrganizationUsers(ctx context.Context, organizationID int64) ([]models.OrganizationUserWithNames, error) {
	var organizationUsers []models.OrganizationUserWithNames

	query := `
        SELECT 
		 	u.first_name,
            u.last_name,
            ou.id AS organization_user_id,
            ou.user_id,
            ou.organization_id,
            ou.role,
            ou.status,
            ou.created_at,
			ou.updated_at          
        FROM 
            organization_users ou
        LEFT JOIN 
            users u ON u.id = ou.user_id
        WHERE 
            ou.organization_id = ?
        ORDER BY 
            ou.updated_at DESC
    `

	// Execute the raw SQL query
	rows, err := t.Client.DB.QueryContext(ctx, query, organizationID)
	if err != nil {
		log.Printf("Error retrieving organization users: %v", err)
		return nil, err
	}
	defer rows.Close()

	// Scan the results into the struct
	for rows.Next() {
		var ou models.OrganizationUserWithNames
		if err := rows.Scan(
			&ou.FirstName,
			&ou.LastName,
			&ou.OrganizationUser.ID,
			&ou.OrganizationUser.UserID,
			&ou.OrganizationUser.OrganizationID,
			&ou.OrganizationUser.Role,
			&ou.OrganizationUser.Status,
			&ou.OrganizationUser.CreatedAt,
			&ou.OrganizationUser.UpdatedAt,
		); err != nil {
			log.Printf("Error scanning row: %v", err)
			return nil, err
		}
		organizationUsers = append(organizationUsers, ou)
	}

	// Check for errors from iterating over rows
	if err := rows.Err(); err != nil {
		log.Printf("Error iterating over rows: %v", err)
		return nil, err
	}

	return organizationUsers, nil
}

func (t DbService) GetOrganizationUserByID(ctx context.Context, id int64) (*models.OrganizationUser, error) {
	organizationUser := new(models.OrganizationUser)
	err := t.Client.DB.NewSelect().Model(organizationUser).Where("id = ?", id).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return organizationUser, nil
}

func (t DbService) GetOrganizationByUserID(ctx context.Context, organizationId, userId int64) (*models.OrganizationUser, error) {
	organizationUser := new(models.OrganizationUser)
	err := t.Client.DB.NewSelect().Model(organizationUser).Where("organization_id = ? AND user_id = ?", organizationId, userId).Scan(ctx)
	if err != nil {
		log.Printf("Error retrieving organizationUser: %v", err)
		return nil, err
	}
	return organizationUser, nil
}

// UpdateOrganizationUser updates an existing OrganizationUser
func (t DbService) UpdateOrganizationUser(ctx context.Context, organizationUser *models.OrganizationUser) error {
	_, err := t.Client.DB.NewUpdate().Model(organizationUser).Where("id = ?", organizationUser.ID).Exec(ctx)
	if err != nil {
		log.Printf("Error updating OrganizationUser: %v", err)
		return err
	}
	return nil
}

// DeleteOrganizationUser deletes a OrganizationUser by ID
func (t DbService) DeleteOrganizationUser(ctx context.Context, id int64) error {
	organizationUser := new(models.OrganizationUser)
	_, err := t.Client.DB.NewDelete().Model(organizationUser).Where("id = ?", id).Exec(ctx)
	if err != nil {
		log.Printf("Error deleting OrganizationUser: %v", err)
		return err
	}
	return nil
}
