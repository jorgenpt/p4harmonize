package p4

import (
	"fmt"
)

// ListDepotFiles runs "p4 fstat" and parses the results into a slice of DepotFile structs.
// Order of resulting slice is alphabetical by Path, ignoring case.
func (p *P4) ListDepotFiles(fileSpecs []string) ([]DepotFile, error) {
	if len(fileSpecs) == 0 {
		fileSpecs = []string{"..."}
	}
	cmd := fmt.Sprintf(`%s fstat -T depotFile,headAction,headChange,headType,digest -Ol `+
		`-F '^(headAction=move/delete | headAction=purge | headAction=archive | headAction=delete)'`, p.cmd())
	for _, fileSpec := range fileSpecs {
		cmd += fmt.Sprintf(" //%s/%s", p.Client, fileSpec)
	}

	return p.runAndParseDepotFiles(cmd)
}
