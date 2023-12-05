package internal

import (
	"embed"
	"os"

	"github.com/devetek/go-core/mdfs"
)

var (
	// Update this embed when create html in different structure

	//go:embed components/**/*.html router/**/view/*.html
	tmpls embed.FS

	Template = mdfs.New(tmpls, "", os.Getenv("ENV"))
)
