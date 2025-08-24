package config

import (
	"fmt"
	"os"
	"strings"
)

func LoadEnv() error {
	env, err := os.ReadFile(".env")
	if err != nil {
		return fmt.Errorf("error reading .env file: %w", err)
	}

	lines := strings.FieldsSeq(string(env))
	for line := range lines {
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			a, b := parts[0], parts[1]
			err := os.Setenv(a, b)
			if err != nil {
				return fmt.Errorf("error setting os env: %w", err)
			}
		}
	}

	return nil
}
