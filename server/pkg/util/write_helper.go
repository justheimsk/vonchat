package util

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/justheimsk/vonchat/server/internal/domain/models"
)

func WriteHTTPError(w http.ResponseWriter, err error) {
	var sErr models.CustomError

	if cErr, ok := err.(*models.CustomError); ok {
		sErr = *cErr
	} else {
		sErr = *models.InternalError
	}

	buf := map[string]interface{}{
		"error": sErr.ToHttpStatusCode(),
	}

	if err := json.NewEncoder(w).Encode(buf); err != nil {
		fmt.Fprintf(w, "Internal server error.")
	}
}
