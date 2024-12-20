package repository

func (r *PostgresRepository) Hesoyam(accountId int) error {
	query := `UPDATE accounts SET balance = balance + 250000 WHERE id = $1`
	_, err := r.db.Exec(query, accountId)

	return err
}
