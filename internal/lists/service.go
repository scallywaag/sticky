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
		formatter.Print(l.Name, l.Id, count, formatter.Default, false)
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

func (s *Service) Delete(id int) error {
	err := s.repo.Delete(id)
	if err != nil {
		return fmt.Errorf("failed to delete list: %w", err)
	}

	err = s.GetAll()
	if err != nil {
		return fmt.Errorf("failed to get lists: %w", err)
	}

	formatter.PrintColored("\nList successfully deleted.", formatter.Yellow)
	return nil
}

func (s *Service) GetActive() (*List, error) {
	l, err := s.repo.GetActive()
	if err != nil {
		return nil, fmt.Errorf("failed to get active list: %w", err)
	}

	return l, nil
}

func (s *Service) SetActive(name string) (*List, error) {
	id, err := s.repo.GetId(name)
	if err != nil {
		return nil, fmt.Errorf("could not get list id: %w", err)
	}

	l, err := s.repo.SetActive(id, name)
	if err != nil {
		return nil, fmt.Errorf("could not set list as active: %w", err)
	}

	return l, nil
}
