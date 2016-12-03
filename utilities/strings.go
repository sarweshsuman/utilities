package utilities

import (
	"bytes"
)

func ConcatStrings(str ...string) string {
	var buffer bytes.Buffer
	for _,st := range str {
		buffer.WriteString(st)
	}
	return buffer.String()
}


