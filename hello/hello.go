package hello

import (
	"fmt"
)

func HelloToName(str string, i int) string {
	return fmt.Sprintf("v3.0.0_modified: Hello to %s, %d", str, i)
}

func test() string {
	sstr := "hien"
	i := 1
	return HelloToName(sstr, i)
}
