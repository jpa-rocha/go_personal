package public

import (
	"embed"
)

//go:embed assets
var Assets embed.FS

//go:embed assets/templates/*
var Templates embed.FS
