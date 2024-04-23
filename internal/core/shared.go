package core

import "strings"

type Myto interface {
	Value() string
}

func Replace(tb Myto, s string, targetKey string) string {
	replaceWith := tb.Value()
	return strings.Replace(s, targetKey, replaceWith, 1)
}
