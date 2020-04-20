package gofunc

import (
	"strings"
)

func ReplaceMacro(str string, macro map[string][]string, macroValue map[string]string) string {

	for k, _mArr := range macro {
		newValue := ""
		if _, ok := macroValue[k]; ok {
			newValue = macroValue[k]
		}
		for _, _macro := range _mArr {
			str = strings.Replace(str, _macro, newValue, -1)
		}
	}

	return str
}


