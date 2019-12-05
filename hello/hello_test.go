package hello

import (
  "testing"
)

func TestHelloToName(t *testing.T) {
  expectResult := "V1: Hello to hien"
  if HelloToName("hien") != expectResult {
    t.Errorf("expect: %s, but result is %s", expectResult, HelloToName("hien"))
  }
}
