package util

import (
	"bytes"
	"fmt"
)

func String2json(fields []string, values []string) []byte {
	if len(fields) == 0 || len(values) == 0{
		fmt.Println("fields 或 values的长度为0")
		return []byte{}
	}

	length := len(fields)
	if len(fields) < len(values) {
		length = len(values)
		for i := 0; i < len(values) - len(fields); i++ {
			fields = append(fields, "")
		}
	} else {
		for i := 0; i < len(fields) - len(values) ; i++ {
			values = append(values, "")
		}
	}

	bf := bytes.NewBuffer([]byte{})
	bf.WriteString("{")

	for i := 0; i < length; i++ {
		bf.WriteString(fmt.Sprintf("\n\t\"%s\":", fields[i]))
		bf.WriteString(fmt.Sprintf("\"%s\",", values[i]))
	}
	bf.WriteString("\n}")
	return bf.Bytes()
}
