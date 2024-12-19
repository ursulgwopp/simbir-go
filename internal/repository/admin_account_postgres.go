package repository

import "github.com/ursulgwopp/simbir-go/internal/models"

// AdminCreateAccount implements service.Repository.
func (r *PostgresRepository) AdminCreateAccount(req models.AdminAccountRequest) (int, error) {
	var id int

	query := `INSERT INTO accounts (username, hash_password, balance, is_admin) VALUES ($1, $2, $3, $4) RETURNING id`
	if err := r.db.QueryRow(query, req.Username, req.Password, req.Balance, req.IsAdmin).Scan(&id); err != nil {
		return -1, err
	}

	return id, nil
}

func (r *PostgresRepository) AdminListAccounts(from int, count int) ([]models.AdminAccountResponse, error) {
	query := `SELECT id, username, balance, is_admin FROM accounts ORDER BY id OFFSET $1 LIMIT $2`
	rows, err := r.db.Query(query, from, count)
	if err != nil {
		return []models.AdminAccountResponse{}, err
	}
	defer rows.Close()

	var accounts []models.AdminAccountResponse
	for rows.Next() {
		var account models.AdminAccountResponse
		if err := rows.Scan(&account.Id, &account.Username, &account.Balance, &account.IsAdmin); err != nil {
			return nil, err
		}

		accounts = append(accounts, account)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return accounts, nil
}

func (r *PostgresRepository) AdminGetAccount(accountId int) (models.AdminAccountResponse, error) {
	var account models.AdminAccountResponse
	account.Id = accountId

	query := `SELECT username, balance, is_admin FROM accounts WHERE id = $1`
	if err := r.db.QueryRow(query, accountId).Scan(&account.Username, &account.Balance, &account.IsAdmin); err != nil {
		return models.AdminAccountResponse{}, err
	}

	return account, nil
}

func (r *PostgresRepository) AdminUpdateAccount(accountId int, req models.AdminAccountRequest) error {
	query := `UPDATE accounts SET username = $1, hash_password = $2, balance = $3, is_admin = $4 WHERE id = $5`
	_, err := r.db.Exec(query, req.Username, req.Password, req.Balance, req.IsAdmin, accountId)

	return err
}

func (r *PostgresRepository) AdminDeleteAccount(accountId int) error {
	query := `DELETE FROM accounts WHERE id = $1`
	_, err := r.db.Exec(query, accountId)

	return err
}
