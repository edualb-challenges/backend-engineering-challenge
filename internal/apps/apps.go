package apps

import (
	"fmt"
)

// NewTreeBabel return a treebabel app with an specific JSON file and window size.
func NewTreeBabel(JSONfile string, wSize uint64) (TreeBabel, error) {
	app := TreeBabel{}

	if JSONfile == "" {
		return app, fmt.Errorf("empty json file not allowed")
	}
	if wSize <= 0 {
		return app, fmt.Errorf("window size must be grather than zero")
	}

	app.TranslationDeliveredFile = JSONfile
	app.WindowSize = wSize
	return app, nil
}
