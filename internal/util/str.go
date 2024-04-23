package util

import (
	"fmt"
	"strings"
)

func GetInner(s string, opens string, closes string) (string, error) {
	firstSplit := strings.Split(s, opens)
	if len(firstSplit) < 2 {
		return "", fmt.Errorf("could not find opens")
	}
	secondSplit := strings.Split(firstSplit[1], closes)
	if len(secondSplit) < 2 {
		return "", fmt.Errorf("could not find closes")
	}
	return secondSplit[0], nil
}

func TrimmedSplit(ln string, sep string) ([]string, error) {
	var out []string = []string{}
	sections := strings.Split(ln, sep)
	for _, sec := range sections {
		out = append(out, strings.TrimSpace(sec))
	}
	if len(out) < 1 {
		return []string{}, fmt.Errorf("could not split")
	}
	return out, nil
}

func TrimmedInnerSplit(s string, opens string, closes string, sep string) ([]string, error) {
	inner, innerErr := GetInner(s, opens, closes)
	if innerErr != nil {
		return []string{}, innerErr
	}
	return TrimmedSplit(inner, sep)
}

func Untab(s string) string {
	return strings.ReplaceAll(s, "\t", " ")
}
