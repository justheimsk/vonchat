package healthCheckService

import "github.com/go-chi/chi/v5"

type Handler interface {
	Load(r chi.Router)
}
