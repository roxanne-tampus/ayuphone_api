package services

import (
	"ayuphone_api/internal/models"

	"context"
	"log"
)

// AssignTechnician assigns a technician to a transaction
func (at DbService) AssignTechnician(ctx context.Context, transactionTechnician *models.TransactionTechnician) error {
	_, err := at.Client.DB.NewInsert().Model(transactionTechnician).Exec(ctx)
	if err != nil {
		log.Printf("Error assigning technician: %v", err)
		return err
	}

	err = at.UpdateTransactionStatus(ctx, transactionTechnician.TransactionID, "accepted")
	if err != nil {
		log.Printf("Error updating status: %v", err)
		return err
	}

	return nil
}

// UnassignTechnician updates the `unassigned_at` field for a specific assignment
func (at DbService) UnassignTechnician(ctx context.Context, id int64) error {
	_, err := at.Client.DB.NewUpdate().Model(&models.TransactionTechnician{}).
		Set("unassigned_at = NOW()").
		Where("id = ?", id).
		Where("unassigned_at IS NULL"). // Only unassign if currently assigned
		Exec(ctx)
	if err != nil {
		log.Printf("Error unassigning technician: %v", err)
		return err
	}
	return nil
}

// GetTechnicianAssignments retrieves the assignment history for a specific transaction
func (at DbService) GetTechnicianAssignments(ctx context.Context, transactionID int64) ([]models.TransactionTechnician, error) {
	var assignments []models.TransactionTechnician
	err := at.Client.DB.NewSelect().
		Model(&assignments).
		Where("transaction_id = ?", transactionID).
		Order("assigned_at ASC").
		Scan(ctx)
	if err != nil {
		log.Printf("Error retrieving assignments: %v", err)
		return nil, err
	}
	return assignments, nil
}
