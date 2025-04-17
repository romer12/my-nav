package assets

import "embed"

var (
	//go:embed static
	Static embed.FS

	//go:embed templates
	Template embed.FS
)
