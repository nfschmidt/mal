package main

import (
	"fmt"
	"strconv"
	"strings"
)

func prStr(mt malType) string {
	switch mt.(type) {
	case []malType:
		str := []string{}
		for _, e := range mt.([]malType) {
			str = append(str, prStr(e))
		}

		return fmt.Sprintf("(%s)", strings.Join(str, " "))
	case malSymbol:
		return mt.(malSymbol).name
	case int:
		return strconv.Itoa(mt.(int))
	}

	return "<INVALID>"
}
