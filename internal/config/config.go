package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/highseas-software/sticky/internal/env"
	"github.com/highseas-software/sticky/internal/formatter"
)

func LoadEnv() error {
	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("cannot get current working directory: %w", err)
	}

	devMode := false
	if _, err := os.Stat(filepath.Join(cwd, "go.mod")); err == nil {
		devMode = true
	}

	if !devMode {
		return nil
	}

	envFile := filepath.Join(cwd, ".env")
	data, err := os.ReadFile(envFile)
	if err != nil {
		return nil
	}

	lines := strings.FieldsSeq(string(data))
	for line := range lines {
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			if err := os.Setenv(parts[0], parts[1]); err != nil {
				return fmt.Errorf("error setting env: %w", err)
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
