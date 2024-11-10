package util

import "fmt"

func ConcatPath(method string, prefix string, endpoint string) string {
  return fmt.Sprintf("%s %s%s", method, prefix, endpoint)
}
