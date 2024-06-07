package utils

import (
	"fmt"
	"strings"
)

func ParseCLIOutput(data string) ([]PasswordEntry, error) {
	var records []PasswordEntry
	lines := strings.Split(strings.TrimSpace(data), "\n")

	for _, line := range lines {
		parts := strings.Split(line, " |\t")

		if len(parts) != 3 {
			return nil, fmt.Errorf("Invalid output format")
		}

		var record PasswordEntry
		for _, part := range parts {
			field := strings.SplitN(part, ": ", 2)
			if len(field) != 2 {
				return nil, fmt.Errorf("Invalid field format")
			}

			switch strings.TrimSpace(field[0]) {
			case "Service":
				record.Service = strings.TrimSpace(field[1])
			case "Username":
				record.Service = strings.TrimSpace(field[1])
			case "Password":
				record.Service = strings.TrimSpace(field[1])
			}
			records = append(records, record)
		}
	}

	return records, nil
}
