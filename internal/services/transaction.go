package services

import (
	"ayuphone_api/internal/models"
	"context"
	"log"
)

// CreateTransaction adds a new transaction to the database
func (t DbService) CreateTransaction(ctx context.Context, transaction *models.Transaction) (*models.Transaction, error) {
	_, err := t.Client.DB.NewInsert().Model(transaction).Returning("*").Exec(ctx)
	if err != nil {
		log.Printf("Error creating transaction: %v", err)
		return nil, err
	}
	return transaction, nil
}

// GetTransaction retrieves all transactions
func (t DbService) GetTransactions(ctx context.Context, userID int64, filter string) ([]models.TransactionWithDevice, error) {
	var transactions []models.TransactionWithDevice
	if userID == 0 {
		if filter == "" {
			err := t.Client.DB.NewSelect().Model(&transactions).Order("updated_at DESC").Scan(ctx)
			if err != nil {
				log.Printf("Error retrieving transactions for user %d: %v", userID, err)
				return nil, err
			}
		} else {
			err := t.Client.DB.NewSelect().Model(&transactions).Where("status = ?", filter).Order("updated_at DESC").Scan(ctx)
			if err != nil {
				log.Printf("Error retrieving transactions for user %d: %v", userID, err)
				return nil, err
			}
		}
	} else {

		rawQuery := `
		SELECT t.id, t.status, t.pickup_address, t.full_address, d.brand, d.model, di.issue_description
		FROM transactions as t
		LEFT JOIN devices AS d ON d.id = t.device_id
		LEFT JOIN device_issues AS di ON di.id = t.device_issue_id
		WHERE t.customer_id = ?
		ORDER BY t.updated_at DESC;`

		err := t.Client.DB.NewRaw(rawQuery, userID).Scan(ctx, &transactions)
		if err != nil {
			log.Printf("Error retrieving transactions for user %d: %v", userID, err)
			return nil, err
		}
	}
	return transactions, nil
}

func (t DbService) GetTransactionByID(ctx context.Context, id int64) (*models.Transaction, error) {
	transaction := new(models.Transaction)
	err := t.Client.DB.NewSelect().Model(transaction).Where("id = ?", id).Scan(ctx)
	if err != nil {
		log.Printf("Error retrieving transaction: %v", err)
		return nil, err
	}
	return transaction, nil
}

// UpdateTransaction updates an existing transaction
func (t DbService) UpdateTransaction(ctx context.Context, transaction *models.Transaction) error {
	_, err := t.Client.DB.NewUpdate().Model(transaction).Where("id = ?", transaction.ID).Exec(ctx)
	if err != nil {
		log.Printf("Error updating transaction: %v", err)
		return err
	}
	return nil
}

func (t DbService) UpdateTransactionStatus(ctx context.Context, transactionID int64, status string) error {
	_, err := t.Client.DB.NewUpdate().Model(&models.Transaction{ID: transactionID, Status: status}).
		Column("status").
		Where("id = ?", transactionID).
		Exec(ctx)
	if err != nil {
		log.Printf("Error updating transaction status: %v", err)
		return err
	}
	return nil
}

// DeleteTransaction deletes a transaction by ID
func (t DbService) DeleteTransaction(ctx context.Context, id int64) error {
	transaction := new(models.Transaction)
	_, err := t.Client.DB.NewDelete().Model(transaction).Where("id = ?", id).Exec(ctx)
	if err != nil {
		log.Printf("Error deleting transaction: %v", err)
		return err
	}
	return nil
}
