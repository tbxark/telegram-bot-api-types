package spec

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/tbxark/telegram-bot-api-types/internal/scrape"
)

func RenderSpec(resp *scrape.APIResponse, dir string) error {
	latestPath := filepath.Join(dir, "spec", "latest.json")

	latest, err := os.Create(latestPath)
	if err != nil {
		return err
	}

	bytes, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		return err
	}

	_, err = latest.Write(bytes)
	if err != nil {
		return err
	}

	return nil
}
