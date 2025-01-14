package repository

import (
	"context"
	"time"

	"github.com/ursulgwopp/simbir-go/internal/models"
)

func (r *PostgresRepository) AdminCreateAccount(req models.AdminAccountRequest) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), operationTimeout)
	defer cancel()

	var id int
	query := `INSERT INTO accounts (username, hash_password, balance, is_admin) VALUES ($1, $2, $3, $4) RETURNING id`
	if err := r.db.QueryRowContext(ctx, query, req.Username, req.Password, req.Balance, req.IsAdmin).Scan(&id); err != nil {
		return -1, err
	}

	return id, nil
}

func (r *PostgresRepository) AdminListAccounts(from int, count int) ([]models.AdminAccountResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	var accounts []models.AdminAccountResponse
	query := `SELECT id, username, balance, is_admin FROM accounts ORDER BY id OFFSET $1 LIMIT $2`
	rows, err := r.db.QueryContext(ctx, query, from, count)
	if err != nil {
		return []models.AdminAccountResponse{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var account models.AdminAccountResponse
		if err := rows.Scan(&account.Id, &account.Username, &account.Balance, &account.IsAdmin); err != nil {
			return []models.AdminAccountResponse{}, err
		}

		accounts = append(accounts, account)
	}

	if err := rows.Err(); err != nil {
		return []models.AdminAccountResponse{}, err
	}

	return accounts, nil
}

func (r *PostgresRepository) AdminGetAccount(accountId int) (models.AdminAccountResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), operationTimeout)
	defer cancel()

	var account models.AdminAccountResponse = models.AdminAccountResponse{Id: accountId}
	query := `SELECT username, balance, is_admin FROM accounts WHERE id = $1`
	if err := r.db.QueryRowContext(ctx, query, accountId).Scan(&account.Username, &account.Balance, &account.IsAdmin); err != nil {
		return models.AdminAccountResponse{}, err
	}

	return account, nil
}

func (r *PostgresRepository) AdminUpdateAccount(accountId int, req models.AdminAccountRequest) error {
	ctx, cancel := context.WithTimeout(context.Background(), operationTimeout)
	defer cancel()

	query := `UPDATE accounts SET username = $1, hash_password = $2, balance = $3, is_admin = $4 WHERE id = $5`
	_, err := r.db.ExecContext(ctx, query, req.Username, req.Password, req.Balance, req.IsAdmin, accountId)

	return err
}

func (r *PostgresRepository) AdminDeleteAccount(accountId int) error {
	ctx, cancel := context.WithTimeout(context.Background(), operationTimeout)
	defer cancel()

	query := `DELETE FROM accounts WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, accountId)

	return err
}
