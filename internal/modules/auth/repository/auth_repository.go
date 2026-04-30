// repository/auth_repository.go
package repository

import "github.com/jmoiron/sqlx"

type Repository struct {
	DB *sqlx.DB
}

type userAuth struct {
	UserID       int64  `db:"id"`
	PasswordHash string `db:"password_hash"`
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

	var ua userAuth

	err := r.DB.Get(&ua, `
		SELECT u.id, ai.password_hash
		FROM users u
		JOIN auth_identities ai ON ai.user_id = u.id
		WHERE u.email = ? AND ai.provider = 'local'
	`, email)

	if err != nil {
		return 0, "", err
	}

	return ua.UserID, ua.PasswordHash, nil
}

// find by provider
func (r *Repository) FindByProvider(provider, pid string) (int64, error) {

	var userID int64

	err := r.DB.Get(&userID, `
		SELECT user_id FROM auth_identities
		WHERE provider = ? AND provider_user_id = ?
	`, provider, pid)

	return userID, err
}
