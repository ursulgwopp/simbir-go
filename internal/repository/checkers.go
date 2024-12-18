package repository

func (r *PostgresRepository) CheckUsernameExists(username string) (bool, error) {
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM accounts WHERE username = $1)"

	err := r.db.QueryRow(query, username).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func (r *PostgresRepository) CheckUsernameIsEqualToOld(accountId int, username string) (bool, error) {
	var oldUsername string

	query := `SELECT username FROM accounts WHERE id = $1`
	if err := r.db.QueryRow(query, accountId).Scan(&oldUsername); err != nil {
		return false, err
	}

	if username != oldUsername {
		return false, nil
	}

	return true, nil
}

func (r *PostgresRepository) CheckTokenIsValid(token string) (bool, error) {
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM blacklist WHERE token = $1)"

	err := r.db.QueryRow(query, token).Scan(&exists)
	if err != nil {
		return false, err
	}

	return !exists, nil
}
