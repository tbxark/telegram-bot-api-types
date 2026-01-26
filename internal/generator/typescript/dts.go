package typescript

import (
	_ "embed"
	"encoding/json"
	"github.com/tbxark/telegram-bot-api-types/internal/generator/tmpl"
	"github.com/tbxark/telegram-bot-api-types/internal/scrape"
	"os"
	"path/filepath"
)

//go:embed dts.tmpl
var dtsTemplate string

func RenderDTS(resp *scrape.APIResponse, dir string) error {
	dtsPath := filepath.Join(dir, "dts", "index.d.ts")
	packageJsonPath := filepath.Join(dir, "dts", "package.json")

	dts, err := os.Create(dtsPath)
	if err != nil {
		return err
	}
	packJson, err := os.Create(packageJsonPath)
	if err != nil {
		return err
	}
	err = tmpl.Render(tmpl.Conf{
		Template: dtsTemplate,
		FuncMap: tmpl.FuncMap{
			UnionsTypes:  UnionsTypes,
			ToFieldTypes: ToFieldTypes,
		},
	}, resp, dts)
	if err != nil {
		return err
	}

	pkg := NewPackage("telegram-bot-api-types", resp.Version)
	pkg.Types = "index.d.ts"
	pkg.Files = []string{"index.d.ts"}

	encoder := json.NewEncoder(packJson)
	encoder.SetIndent("", "  ")
	return encoder.Encode(pkg)
}
