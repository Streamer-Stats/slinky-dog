package controllers
import (
	"github.com/go-openapi/runtime/middleware"
	oapi"leagueapi.com.br/rest/restapi/operations/api"
)
// PlayerController controller of player acts
type PlayerController struct {}

// GetDamage get damage from player
func (pc *PlayerController) GetDamage(params oapi.LiveMatchParams) middleware.Responder {
	return oapi.NewLiveMatchOK()
}

// NewPlayerController Constructor playerController
func NewPlayerController() *PlayerController{
	return &PlayerController{}
}