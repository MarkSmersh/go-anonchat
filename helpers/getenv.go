package helpers

import (
	"os"
	"strings"
)

func GetEnv() (map[string]string, error) {
	file, err := os.ReadFile(".env")

	if err != nil {
		return nil, err
	}

	lines := map[string]string{}

	for _, line := range strings.Split(string(file[:]), "\n") {
		splittedLine := strings.Split(line, "=")

		if splittedLine[0] != "" && splittedLine[1] != "" {
			lines[splittedLine[0]] = splittedLine[1]
		}

	}

	return lines, nil
}
