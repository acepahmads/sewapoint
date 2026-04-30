// usecase/auth_usecase.go
package usecase

import (
	"errors"

	"sewapoint/internal/modules/auth/dto"
	"sewapoint/internal/modules/auth/repository"
	"sewapoint/pkg/jwt"

	"golang.org/x/crypto/bcrypt"
	"google.golang.org/api/idtoken"
)

type Usecase struct {
	Repo *repository.Repository
}

func (u *Usecase) Register(req dto.RegisterRequest) (string, error) {

	hash, _ := bcrypt.GenerateFromPassword([]byte(req.Password), 10)

	userID, err := u.Repo.CreateUser(req.Name, req.Email)
	if err != nil {
		return "", err
	}

	err = u.Repo.CreateIdentity(userID, "local", req.Email, string(hash))
	if err != nil {
		return "", err
	}

	return jwt.GenerateToken(int(userID))
}

func (u *Usecase) Login(req dto.LoginRequest) (string, error) {

	userID, hash, err := u.Repo.FindUserByEmail(req.Email)
	if err != nil {
		return "", errors.New("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(req.Password))
	if err != nil {
		return "", errors.New("invalid password")
	}

	return jwt.GenerateToken(int(userID))
}

func (u *Usecase) SocialLogin(req dto.SocialLoginRequest) (string, error) {

	// verify google token
	payload, err := idtoken.Validate(nil, req.Token, "")
	if err != nil {
		return "", errors.New("invalid token")
	}

	providerID := payload.Subject
	email := payload.Claims["email"].(string)
	name := payload.Claims["name"].(string)

	userID, err := u.Repo.FindByProvider(req.Provider, providerID)

	if err == nil {
		return jwt.GenerateToken(int(userID))
	}

	// create new user
	userID, err = u.Repo.CreateUser(name, email)
	if err != nil {
		return "", err
	}

	err = u.Repo.CreateIdentity(userID, req.Provider, providerID, "")
	if err != nil {
		return "", err
	}

	return jwt.GenerateToken(int(userID))
}
