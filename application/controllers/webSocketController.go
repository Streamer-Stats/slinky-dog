package controllers
import (
	"leagueapi.com.br/rest/pkg/infrastructure/socket"
	"github.com/go-openapi/runtime/middleware"
	oapi"leagueapi.com.br/rest/restapi/operations/live"
	"net/http"
	"github.com/go-openapi/runtime"
)
// WebSocketController controller of WebSocket acts
type WebSocketController struct {
	socket *socket.Socket
}

// Connect get damage from WebSocket
func (wsc *WebSocketController) Connect(params oapi.LiveParams) middleware.Responder {
	return middleware.ResponderFunc(func(rw http.ResponseWriter, _ runtime.Producer) {
		wsc.socket.CreateConn(rw, params.HTTPRequest).Handle()
	})
}

// NewWebSocketController Constructor WebSocketController
func NewWebSocketController(_socket *socket.Socket) *WebSocketController{
	return &WebSocketController{
		socket: _socket,
	}
}