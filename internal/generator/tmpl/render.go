package tmpl

import (
	"github.com/tbxark/telegram-bot-api-types/internal/scrape"
	"io"
	"strings"
	"text/template"
)

type FuncMap struct {
	UnionsTypes  func(types []string) string
	ToFieldTypes func(*scrape.Field) string
	ExtraFunc    map[string]any
}

type Conf struct {
	Template string
	FuncMap  FuncMap
}

func HasResponse(types []string) bool {
	if len(types) == 0 {
		return false
	}
	if len(types) == 1 && types[0] == "Boolean" {
		return false
	}
	return true
}

func HasParams(method *scrape.Method) bool {
	return len(method.Fields) > 0
}

func IsAbstractType(t *scrape.Type) bool {
	return len(t.Fields) == 0 && len(t.Subtypes) > 0
}

func IsParamsOptional(fields []*scrape.Field) bool {
	for _, f := range fields {
		if f.Required {
			return false
		}
	}
	return true
}

func ToPascalCase(str string) string {
	return strings.ToUpper(string(str[0])) + str[1:]
}

func ToCamelCase(str string) string {
	return strings.ToLower(string(str[0])) + str[1:]
}

func ToTypesDoc(types []string) string {
	return strings.Join(types, " or ")
}

func IsKeyword(words []string) func(word string) bool {
	wordSet := make(map[string]struct{}, len(words))
	for _, w := range words {
		wordSet[w] = struct{}{}
	}
	return func(word string) bool {
		_, ok := wordSet[word]
		return ok
	}
}

func Render(conf Conf, resp *scrape.APIResponse, writer io.Writer) error {
	funcMap := template.FuncMap{
		"UnionsTypes":      conf.FuncMap.UnionsTypes,
		"ToFieldTypes":     conf.FuncMap.ToFieldTypes,
		"ToPascalCase":     ToPascalCase,
		"ToCamelCase":      ToCamelCase,
		"ToTypesDoc":       ToTypesDoc,
		"HasResponse":      HasResponse,
		"HasParams":        HasParams,
		"IsAbstractType":   IsAbstractType,
		"IsParamsOptional": IsParamsOptional,
	}
	if conf.FuncMap.ExtraFunc != nil {
		for k, v := range conf.FuncMap.ExtraFunc {
			funcMap[k] = v
		}
	}
	tmpl, err := template.New("tmpl").Funcs(funcMap).Parse(conf.Template)
	if err != nil {
		return err
	}
	err = tmpl.Execute(writer, resp)
	if err != nil {
		return err
	}
	return nil
}
