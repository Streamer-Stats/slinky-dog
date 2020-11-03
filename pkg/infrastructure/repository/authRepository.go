package repository

import (
	"github.com/go-pg/pg/v10"
	"leagueapi.com.br/rest/models"
	"leagueapi.com.br/rest/pkg/domain/entities"
	"leagueapi.com.br/rest/pkg/infrastructure/database"
)

type AuthRepository struct {
	DB *pg.DB
}

func (repository *AuthRepository) GetUser(user *models.User) *models.User {
	userPG := &entities.User{}
	err := repository.DB.Model(userPG).Where(`"Email" = ? AND "Password" = ?`, user.Email, user.Password).Select()
	if err != nil {
		return nil
	}

	user.Username = userPG.Username
	user.ID = userPG.ID
	user.Password = "noshow"
	return user
}

func NewAuthRepository(_database *database.Database) *AuthRepository {
	return &AuthRepository{
		DB: _database.GetConnection(),
	}
}
