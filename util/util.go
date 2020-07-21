package util

import (
	"bytes"
	"fmt"
)

type LengthError struct {
	info string
}

func (err LengthError) Error() string {
	return err.info
}

func String2json(fields []string, values []string) ([]byte, error) {
	if len(fields) == 0 || len(values) == 0{
		return []byte{}, LengthError{"fields 或 values的长度为0"}
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
	//fmt.Println("field len:", len(fields))
	//fmt.Println("values len:", len(values))
	//fmt.Println("length len:", length)

	for i := 0; i < length; i++ {
		bf.WriteString(fmt.Sprintf("\n\t\"%s\":", fields[i]))
		bf.WriteString(fmt.Sprintf("\"%s\",", values[i]))
	}
	bf.WriteString("\n}")
	return bf.Bytes(), nil
}
