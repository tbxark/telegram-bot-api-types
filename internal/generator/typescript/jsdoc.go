package typescript

import (
	_ "embed"
	"encoding/json"
	"github.com/tbxark/telegram-bot-api-types/internal/generator/tmpl"
	"github.com/tbxark/telegram-bot-api-types/internal/scrape"
	"os"
	"path/filepath"
)

//go:embed jsdoc.tmpl
var jsDocTemplate string

func RenderJsDoc(resp *scrape.APIResponse, dir string) error {

	jsDocPath := filepath.Join(dir, "jsdoc", "index.js")
	packageJsonPath := filepath.Join(dir, "jsdoc", "package.json")

	jsDoc, err := os.Create(jsDocPath)
	if err != nil {
		return err
	}
	packJson, err := os.Create(packageJsonPath)
	if err != nil {
		return err
	}

	err = tmpl.Render(tmpl.Conf{
		Template: jsDocTemplate,
		FuncMap: tmpl.FuncMap{
			UnionsTypes:  UnionsTypes,
			ToFieldTypes: ToFieldTypes,
		},
	}, resp, jsDoc)
	if err != nil {
		return err
	}

	pkg := NewPackage("telegram-bot-api-jsdoc", resp.Version)
	pkg.Main = "index.js"
	pkg.Module = "index.js"
	pkg.Files = []string{"index.js"}

	encoder := json.NewEncoder(packJson)
	encoder.SetIndent("", "  ")
	return encoder.Encode(pkg)
}
