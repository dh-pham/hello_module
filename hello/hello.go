package hello

import (
  "fmt"
)

func HelloToName(str string) string {
  return fmt.Sprintf("Hello to %s", str)
}
