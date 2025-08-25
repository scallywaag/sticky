package lists

import (
	"fmt"

	"github.com/highseas-software/sticky/internal/formatter"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetAll() error {
	count, err := s.repo.Count()
	if err != nil {
		return fmt.Errorf("count failed: %w", err)
	}

	lists, err := s.repo.GetAll()
	if err != nil {
		return fmt.Errorf("failed to get lists: %w", err)
	}

	formatter.ClearScreen()
	formatter.PrintListHeader("lists", count)

	for _, l := range lists {
		formatter.Print(l.Name, l.Id, count, string(formatter.Default), false)
	}

	return nil
}

func (s *Service) Add(name string) error {
	_, err := s.repo.Add(name)
	if err != nil {
		return fmt.Errorf("failed to add list: %w", err)
	}

	err = s.GetAll()
	if err != nil {
		return fmt.Errorf("failed to get lists: %w", err)
	}

	formatter.PrintColored("\nList successfully added.", formatter.Yellow)
	return nil
}
