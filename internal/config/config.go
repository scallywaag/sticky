package config

import (
	"os"

	"github.com/scallywaag/sticky/internal/env"
	"github.com/scallywaag/sticky/internal/formatter"
)

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
