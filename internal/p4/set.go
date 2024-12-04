package p4

import (
	"fmt"
	"strings"
)

// GetVariable fetches the local Perforce setting named variable
func (p *P4) GetVariable(variable string) (string, error) {
	var variableValue string = ""
	err := p.cmdAndScan(
		fmt.Sprintf("%s set -q %s", p.cmd(), variable),
		func(rawLine string) error {
			value, found := strings.CutPrefix(rawLine, fmt.Sprintf("%s=", variable))
			if found {
				variableValue = value
			}
			return nil
		},
	)

	if err != nil {
		return "", err
	}

	return variableValue, nil
}
