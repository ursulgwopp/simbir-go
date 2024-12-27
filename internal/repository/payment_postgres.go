package repository

import (
	"context"
	"time"
)

func (r *PostgresRepository) Hesoyam(accountId int) error {
	ctx, cancel := context.WithTimeout(context.Background(), operationTimeout)
	defer cancel()

	query := `UPDATE accounts SET balance = balance + 100000 WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, accountId)

	return err
}

type WithdrawInfo struct {
	RentId      int
	UserId      int
	PriceOfUnit int
}

func (r *PostgresRepository) MinutelyPayment() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	var users []WithdrawInfo
	query := `SELECT id, user_id, price_of_unit FROM rents WHERE is_active = TRUE AND price_type = 'Minutes'`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var user WithdrawInfo
		if err := rows.Scan(&user.RentId, &user.UserId, &user.PriceOfUnit); err != nil {
			return err
		}

		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return err
	}

	for _, user := range users {
		var balance int
		query := `SELECT balance FROM accounts WHERE id = $1`
		if err := r.db.QueryRowContext(ctx, query, user.UserId).Scan(&balance); err != nil {
			return err
		}

		if balance < user.PriceOfUnit {
			r.StopRent(user.RentId, 61, 31)
			continue
		}

		query = `UPDATE accounts SET balance = balance - $1 WHERE id = $2`
		_, err := r.db.ExecContext(ctx, query, user.PriceOfUnit, user.UserId)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *PostgresRepository) DailyPayment() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	var users []WithdrawInfo
	query := `SELECT id, user_id, price_of_unit FROM rents WHERE is_active = TRUE AND price_type = 'Days'`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var user WithdrawInfo
		if err := rows.Scan(&user.RentId, &user.UserId, &user.PriceOfUnit); err != nil {
			return err
		}

		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return err
	}

	for _, user := range users {
		var balance int
		query := `SELECT balance FROM accounts WHERE id = $1`
		if err := r.db.QueryRowContext(ctx, query, user.UserId).Scan(&balance); err != nil {
			return err
		}

		if balance < user.PriceOfUnit {
			r.StopRent(user.RentId, 61, 31)
			continue
		}

		query = `UPDATE accounts SET balance = balance - $1 WHERE id = $2`
		_, err := r.db.ExecContext(ctx, query, user.PriceOfUnit, user.UserId)
		if err != nil {
			return err
		}
	}

	return nil
}
