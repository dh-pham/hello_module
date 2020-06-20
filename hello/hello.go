package hello

import (
	"fmt"
)

func HelloToName(str string, i int) string {
	return fmt.Sprintf("V2_update: Hello to %s, %d", str, i)
}

func test() string {
	sstr := "hien"
	i := 1
	return HelloToName(sstr, i)
}
