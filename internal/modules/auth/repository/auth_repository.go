// repository/auth_repository.go
package repository

import (
	"database/sql"
)

type Repository struct {
	DB *sql.DB
}

// create user
func (r *Repository) CreateUser(name, email string) (int64, error) {
	res, err := r.DB.Exec(`
		INSERT INTO users (name, email) VALUES (?, ?)
	`, name, email)

	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}

// create identity
func (r *Repository) CreateIdentity(userID int64, provider, pid, password string) error {
	_, err := r.DB.Exec(`
		INSERT INTO auth_identities (user_id, provider, provider_user_id, password_hash)
		VALUES (?, ?, ?, ?)
	`, userID, provider, pid, password)

	return err
}

// find user by email
func (r *Repository) FindUserByEmail(email string) (int64, string, error) {
	var userID int64
	var password string

	err := r.DB.QueryRow(`
		SELECT u.id, ai.password_hash
		FROM users u
		JOIN auth_identities ai ON ai.user_id = u.id
		WHERE u.email = ? AND ai.provider = 'local'
	`, email).Scan(&userID, &password)

	return userID, password, err
}

// find by provider
func (r *Repository) FindByProvider(provider, pid string) (int64, error) {
	var userID int64

	err := r.DB.QueryRow(`
		SELECT user_id FROM auth_identities
		WHERE provider = ? AND provider_user_id = ?
	`, provider, pid).Scan(&userID)

	return userID, err
}
