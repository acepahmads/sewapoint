// internal/modules/auth/entity/auth_identity.go
package entity

type AuthIdentity struct {
	ID             int64
	UserID         int64
	Provider       string // google, facebook, local
	ProviderUserID string
	PasswordHash   string
}
