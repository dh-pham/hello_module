package hello

import (
  "testing"
)

func TestHelloToName(t *testing.T) {
  expectResult := "v3.0.0: Hello to hien, 1"
  if HelloToName("hien", 1) != expectResult {
    t.Errorf("expect: %s, but result is %s", expectResult, HelloToName("hien", 1))
  }
}
