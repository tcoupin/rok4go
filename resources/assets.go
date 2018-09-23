// +build !dist

//go:generate go run assets_generate.go

package resources

import "net/http"

// Assets contains project assets.
var Assets http.FileSystem = http.Dir("resources/ui/dist")
