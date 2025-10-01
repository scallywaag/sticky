package database

import (
	"log"
	"os"
	"path/filepath"

	"github.com/scallywaag/sticky/internal/env"
)

func getDbPath(appEnv env.StickyEnv) string {
	var dbDir, filename string

	switch appEnv {
	case env.EnvProd:
		homeDir, err := os.UserHomeDir()
		if err != nil {
			log.Fatal(err)
		}
		dbDir = filepath.Join(homeDir, ".local", "share", "sticky")
		filename = "sticky.db"

	case env.EnvDev:
		cwd, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		dbDir = cwd
		filename = "sticky.db"
	}

	dbPath := filepath.Join(dbDir, filename)

	if err := os.MkdirAll(filepath.Dir(dbPath), 0700); err != nil {
		log.Fatalf("Error creating directory: %v", err)
	}

	return dbPath
}

func fileExists(path string) bool {
	info, err := os.Stat(path)
	return err == nil && !info.IsDir()
}
