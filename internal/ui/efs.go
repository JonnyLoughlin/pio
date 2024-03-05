package ui

import "embed"

//go:embed "src/assets"
var AssetFiles embed.FS

//go:embed "src/css"
var CssFiles embed.FS

//go:embed "src/html"
var HtmlFiles embed.FS

//go:embed "src/js"
var JsFiles embed.FS
