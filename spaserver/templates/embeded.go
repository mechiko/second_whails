package templates

import (
	"embed"
)

//go:embed header footer home root kmstate adjust cisinfo menu money target gtin
var root embed.FS
