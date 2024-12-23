package repository

import "context"

func (r *PostgresRepository) Hesoyam(accountId int) error {
	ctx, cancel := context.WithTimeout(context.Background(), operationTimeout)
	defer cancel()

	query := `UPDATE accounts SET balance = balance + 100000 WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, accountId)

	return err
}
