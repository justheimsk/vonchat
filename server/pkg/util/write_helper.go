package util

import (
  "encoding/json"
  "fmt"
  "net/http"

  "github.com/justheimsk/vonchat/server/internal/domain/models"
)

func WriteHTTPError(w http.ResponseWriter, err error) {
  var buf map[string]interface{}

  if cErr, ok := err.(*models.CustomError); ok {
    buf = map[string]interface{}{
      "error": cErr.ToHttpStatusCode(),
    }
  } else if mErr, ok := err.(*models.MultiError); ok { 
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
