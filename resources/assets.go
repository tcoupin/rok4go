// +build dev generate test

//go:generate go run -tags=dev assets_generate.go

package resources

import "net/http"

// Assets contains project assets.
var Assets http.FileSystem = http.Dir("resources/ui/dist")
