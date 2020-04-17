package gofunc

import (
	"strings"
)

type Macro map[string][]string
type MacroValue map[string]string

func ReplaceMacro(str string, macro Macro, macroValue MacroValue) string {

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


