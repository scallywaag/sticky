package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/highseas-software/sticky/internal/env"
	"github.com/highseas-software/sticky/internal/formatter"
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

	PrintAppEnv()

	return nil
}

func GetAppEnv() env.StickyEnv {
	val, ok := os.LookupEnv(env.StickyEnvVar)
	if !ok || val == "" { // default to prod
		val = string(env.EnvProd)
	}

	return env.StickyEnv(val)
}

func PrintAppEnv() {
	val := GetAppEnv()

	switch val {
	case env.EnvDev:
		formatter.PrintColored("Running in dev mode", formatter.Yellow)
	case env.EnvTest:
		formatter.PrintColored("Running in test mode", formatter.Yellow)
	}
}
