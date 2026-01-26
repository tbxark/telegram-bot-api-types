package typescript

import (
	"fmt"
	"github.com/tbxark/telegram-bot-api-types/internal/scrape"
	"strings"
)

func toTypeScriptType(t string) string {
	switch t {
	case "Integer", "Float":
		return "number"
	case "String":
		return "string"
	case "Boolean":
		return "boolean"
	default:
		return t
	}
}

func UnionsTypes(types []string) string {
	if len(types) == 0 {
		return "never"
	}
	either := make([]string, 0, len(types))
	for _, t := range types {
		arrayWrap := 0
		_type := t
		for strings.HasPrefix(_type, "Array of ") {
			_type = _type[len("Array of "):]
			arrayWrap++
		}
		if arrayWrap == 0 {
			either = append(either, toTypeScriptType(_type))
		} else {
			var sb strings.Builder
			for i := 0; i < arrayWrap; i++ {
				sb.WriteString("Array<")
			}
			sb.WriteString(toTypeScriptType(_type))
			for i := 0; i < arrayWrap; i++ {
				sb.WriteString(">")
			}
			either = append(either, sb.String())
		}
	}
	return strings.Join(either, " | ")
}

func ToFieldTypes(field *scrape.Field) string {
	if field.Const != "" && len(field.Types) == 1 && field.Types[0] == "String" {
		return fmt.Sprintf("'%s'", field.Const)
	}
	return UnionsTypes(field.Types)
}
