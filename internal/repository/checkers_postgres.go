package repository

func (r *PostgresRepository) CheckUsernameExists(username string) (bool, error) {
	var exists bool

	query := `SELECT EXISTS(SELECT 1 FROM accounts WHERE username = $1)`
	if err := r.db.QueryRow(query, username).Scan(&exists); err != nil {
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

	query := `SELECT EXISTS(SELECT 1 FROM blacklist WHERE token = $1)`
	if err := r.db.QueryRow(query, token).Scan(&exists); err != nil {
		return false, err
	}

	return !exists, nil
}

func (r *PostgresRepository) CheckAccountIdExists(accountId int) (bool, error) {
	var exists bool

	query := `SELECT EXISTS(SELECT 1 FROM accounts WHERE id = $1)`
	if err := r.db.QueryRow(query, accountId).Scan(&exists); err != nil {
		return false, err
	}

	return exists, nil
}

func (r *PostgresRepository) CheckOwnerId(transportId int) (int, error) {
	var id int

	query := `SELECT owner_id FROM transports WHERE id = $1`
	if err := r.db.QueryRow(query, transportId).Scan(&id); err != nil {
		return -1, err
	}

	return id, nil
}

func (r *PostgresRepository) CheckTransportIdExists(transportId int) (bool, error) {
	var exists bool

	query := `SELECT EXISTS(SELECT 1 FROM transports WHERE id = $1)`
	if err := r.db.QueryRow(query, transportId).Scan(&exists); err != nil {
		return false, err
	}

	return exists, nil
}

func (r *PostgresRepository) CheckRentIdExists(rentId int) (bool, error) {
	var exists bool

	query := `SELECT EXISTS(SELECT 1 FROM rents WHERE id = $1)`
	if err := r.db.QueryRow(query, rentId).Scan(&exists); err != nil {
		return false, err
	}

	return exists, nil
}

func (r *PostgresRepository) CheckRentIsActive(rentId int) (bool, error) {
	var active bool

	query := `SELECT is_active FROM rents WHERE id = $1`
	if err := r.db.QueryRow(query, rentId).Scan(&active); err != nil {
		return false, err
	}

	return active, nil
}

func (r *PostgresRepository) CheckTransportIsAvailable(transportId int) (bool, error) {
	var available bool

	query := `SELECT can_be_rented FROM transports WHERE id = $1`
	if err := r.db.QueryRow(query, transportId).Scan(&available); err != nil {
		return false, err
	}

	return available, nil
}

func (r *PostgresRepository) CheckRentOwnerId(rentId int) (int, error) {
	var id int

	query := `SELECT user_id FROM rents WHERE id = $1`
	if err := r.db.QueryRow(query, rentId).Scan(&id); err != nil {
		return -1, err
	}

	return id, nil
}
