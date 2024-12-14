package ws_delivery

import (
	"github.com/justheimsk/vonchat/server/internal/domain/models"
	"github.com/justheimsk/vonchat/server/internal/domain/payload"
	domain_service "github.com/justheimsk/vonchat/server/internal/domain/service"
	"github.com/justheimsk/vonchat/server/internal/infra/ws"
	"github.com/mitchellh/mapstructure"
)

type IdentifyHandler struct {
	authService domain_service.AuthService
	userService domain_service.UserService
	logger      models.Logger
}

func NewIdentifyHandler(authService domain_service.AuthService, userService domain_service.UserService, logger models.Logger) *IdentifyHandler {
	return &IdentifyHandler{
		authService: authService,
		userService: userService,
		logger:      logger,
	}
}

func (self *IdentifyHandler) Handle(w *ws.WebsocketWriter) {
	if w.Client.Authenticated {
		return
	}

	var data payload.IdentifyPayload
	mapstructure.Decode(w.Message.Data, &data)

	if data.Token == "" {
		w.CloseSocket()
		return
	}

	claims, err := self.authService.ValidateToken(data.Token)
	if err != nil {
		w.CloseSocket()
		return
	}

	id, err := self.authService.GetIdFromClaims(claims)
	if err != nil {
		w.CloseSocket()
		return
	}

	user, err := self.userService.GetUserById(id)
	if err != nil {
		w.CloseSocket()
		return
	}

	w.Client.Authenticate(user)
	w.Write(1, "READY", nil)
	self.userService.SetUserStatus(user.ID, "online")
	self.logger.Infof("Client %s with username \"%s\" is now authenticated", w.Client.RandomID, user.Username)
}
