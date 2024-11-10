package util

import (
  "strings"
)

func JoinErrors(errs []error) string {
  var msgs []string

  for _, err := range(errs) {
    msgs = append(msgs, err.Error())
  }

  return strings.Join(msgs, ", ")
}
