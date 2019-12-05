package hello

import (
  "fmt"
)

func HelloToName(str string) string {
  return fmt.Sprintf("V1: Hello to %s", str)
}
