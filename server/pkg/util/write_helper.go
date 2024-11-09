package util

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/justheimsk/vonchat/server/internal/domain/models"
)

func WriteHTTPError(w http.ResponseWriter, err error) {
	var sErr models.CustomError
  var mErr models.MultiError
  var errType string

	if cErr, ok := err.(*models.CustomError); ok {
		sErr = *cErr
    errType = "single"
	} else if mmErr, ok := err.(*models.MultiError); ok { 
    mErr = *mmErr
    errType = "multi"
  } else {
		sErr = *models.InternalError
	}

  var buf map[string]interface{}
  if errType == "single" {
    code, _ := strconv.Atoi(sErr.ToHttpStatusCode().Code)
    w.WriteHeader(code)

	  buf = map[string]interface{}{
		  "error": sErr.ToHttpStatusCode(),
	  }
  } else if errType == "multi" {
    code, _ := strconv.Atoi(mErr.Code)
    w.WriteHeader(code)

    buf = map[string]interface{}{
      "code": mErr.Code,
      "errors": mErr.Errors,
    }
  }

	if err := json.NewEncoder(w).Encode(buf); err != nil {
		fmt.Fprintf(w, "Internal server error.")
	}
}

func WriteHTTPResponse(w http.ResponseWriter, buf map[string]interface{}) {
	if err := json.NewEncoder(w).Encode(buf); err != nil {
		WriteHTTPError(w, models.InternalError)
	}
}
