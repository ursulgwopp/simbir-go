package repository

import (
	"github.com/ursulgwopp/simbir-go/internal/models"
)

func (r *PostgresRepository) SignUp(req models.AccountRequest) (int, error) {
	var id int

	query := `INSERT INTO accounts (username, hash_password) VALUES ($1, $2) RETURNING id`
	if err := r.db.QueryRow(query, req.Username, req.Password).Scan(&id); err != nil {
		return -1, err
	}

	return id, nil
}

func (r *PostgresRepository) SignIn(req models.AccountRequest) (models.TokenInfo, error) {
	var id int
	var isAdmin bool

	query := `SELECT id, is_admin FROM accounts WHERE username = $1 and hash_password = $2`
	if err := r.db.QueryRow(query, req.Username, req.Password).Scan(&id, &isAdmin); err != nil {
		return models.TokenInfo{}, err
	}
	// row := r.db.QueryRow(query, req.Username, req.Password)
	// if err := row.Scan(&id, &isAdmin); err != nil {
	// 	return models.TokenInfo{}, err
	// }

	return models.TokenInfo{AccountId: id, IsAdmin: isAdmin}, nil
}

func (r *PostgresRepository) SignOut(token string) error {
	query := `INSERT INTO blacklist (token) VALUES ($1)`
	_, err := r.db.Exec(query, token)

	return err
}

func (r *PostgresRepository) GetAccount(accountId int) (models.AccountResponse, error) {
	var account models.AccountResponse
	account.Id = accountId

	query := `SELECT username, balance FROM accounts WHERE id = $1`
	if err := r.db.QueryRow(query, accountId).Scan(&account.Username, &account.Balance); err != nil {
		return models.AccountResponse{}, err
	}

	return account, nil
}

func (r *PostgresRepository) UpdateAccount(accountId int, req models.AccountRequest) error {
	query := `UPDATE accounts SET username = $1, hash_password = $2 WHERE id = $3`
	_, err := r.db.Exec(query, req.Username, req.Password, accountId)

	return err
}
