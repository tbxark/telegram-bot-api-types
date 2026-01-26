package main

import (
	"flag"
	"github.com/tbxark/telegram-bot-api-types/internal/generator/spec"
	"github.com/tbxark/telegram-bot-api-types/internal/generator/swift"
	"github.com/tbxark/telegram-bot-api-types/internal/generator/typescript"
	"github.com/tbxark/telegram-bot-api-types/internal/scrape"
	"log"
	"strings"
)

type Generator interface {
	Generate(resp *scrape.APIResponse, dir string) error
}

type GeneratorFunc func(resp *scrape.APIResponse, dir string) error

func (f GeneratorFunc) Generate(resp *scrape.APIResponse, dir string) error {
	return f(resp, dir)
}

var generators = map[string]Generator{
	"typescript": GeneratorFunc(typescript.RenderDTS),
	"jsdoc":      GeneratorFunc(typescript.RenderJsDoc),
	"spec":       GeneratorFunc(spec.RenderSpec),
	"swift":      GeneratorFunc(swift.RenderSwift),
}

func main() {

	dist := flag.String("dist", "./dist", "The output directory")
	lang := flag.String("lang", "typescript,jsdoc,spec,swift", "The output language")
	help := flag.Bool("help", false, "Show help")

	flag.Parse()
	if *help {
		flag.PrintDefaults()
		return
	}

	items, err := scrape.RetrieveInfo()
	if err != nil {
		log.Fatalf("Failed to retrieve info: %v", err)
	}
	if scrape.Verify(items) {
		log.Fatalf("Errors found in API data")
	}
	for _, l := range strings.Split(*lang, ",") {
		gen, ok := generators[l]
		if !ok {
			log.Fatalf("Unknown language: %s", l)
		}
		e := gen.Generate(items, *dist)
		if e != nil {
			log.Fatalf("Failed to generate: %v", e)
		}
	}
}
