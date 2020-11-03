package services

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"leagueapi.com.br/rest/models"
	"leagueapi.com.br/rest/pkg/helpers"
	"leagueapi.com.br/rest/pkg/infrastructure/repository"
)

type AuthService struct {
	AuthRepository *repository.AuthRepository
}

func (service *AuthService) checkUserEmpty(user *models.User) bool {
	if user != nil {
		return user.Email != "" && user.Password != ""
	}
	return false

}

func (service *AuthService) Login(user *models.User) []*models.User {
	var users []*models.User

	if service.checkUserEmpty(user) {
		fetchedUser := service.AuthRepository.GetUser(user)
		if service.checkUserEmpty(fetchedUser) {
			token := jwt.New(jwt.SigningMethodHS256)
			claims := token.Claims.(jwt.MapClaims)

			claims["UserID"] = fetchedUser.ID
			claims["UserName"] = user.Username
			claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

			tokenString, err := token.SignedString([]byte(helpers.MySigningKey))

			if err != nil {
				panic(err)
			}

			fetchedUser.Token = tokenString
			users = append(users, fetchedUser)
			return users
		}
	}

	return nil
}

func NewAuthService(_authRepository *repository.AuthRepository) *AuthService {
	return &AuthService{
		AuthRepository: _authRepository,
	}
}
