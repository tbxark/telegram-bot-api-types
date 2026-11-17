package typescript

import "strings"

// Repository uses the https URL on the canonical lowercase owner: npm's provenance
// check rejects the ssh form and the mixed-case owner used by the git remote.
type Repository struct {
	Type string `json:"type"`
	URL  string `json:"url"`
}

type Package struct {
	Name        string     `json:"name"`
	Version     string     `json:"version"`
	Description string     `json:"description"`
	Repository  Repository `json:"repository"`
	Type        string     `json:"type"`
	Author      string     `json:"author"`
	License     string     `json:"license"`
	Types       string     `json:"types,omitempty"`
	Main        string     `json:"main,omitempty"`
	Module      string     `json:"module,omitempty"`
	Files       []string   `json:"files,omitempty"`
}

func NewPackage(name, version string) *Package {
	if strings.Count(version, ".") < 2 {
		version += ".0"
	}
	return &Package{
		Name:        name,
		Version:     version,
		Description: "Telegram Bot API types",
		Repository: Repository{
			Type: "git",
			URL:  "git+https://github.com/tbxark/telegram-bot-api-types.git",
		},
		Type:    "module",
		Author:  "TBXark",
		License: "MIT",
	}
}
