package ws_delivery

import (
	"github.com/justheimsk/vonchat/server/internal/domain/models"
	domain_service "github.com/justheimsk/vonchat/server/internal/domain/service"
	"github.com/justheimsk/vonchat/server/internal/infra/ws"
	"github.com/mitchellh/mapstructure"
)

type IdentifyHandler struct {
	authService domain_service.AuthService
	userService domain_service.UserService
	logger      models.Logger
}

type IdentifyPayload struct {
	Token string `json:"token"`
}

func NewIdentifyHandler(authService domain_service.AuthService, userService domain_service.UserService, logger models.Logger) *IdentifyHandler {
	return &IdentifyHandler{
		authService: authService,
		userService: userService,
		logger:      logger,
	}
}

func (self *IdentifyHandler) Handle(w *ws.WebsocketWriter) {
	var data IdentifyPayload
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
	self.logger.Infof("Client %s with username \"%s\" is now authenticated", w.Client.RandomID, user.Username)
}
