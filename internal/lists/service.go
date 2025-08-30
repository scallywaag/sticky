package lists

import (
	"errors"
	"fmt"

	"github.com/highseas-software/sticky/internal/formatter"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

var (
	clearScreen     = formatter.ClearScreen
	printListHeader = formatter.PrintListHeader
	printContent    = formatter.PrintContent
	printColored    = formatter.PrintColored
)

func (s *Service) GetAll() error {
	count, err := s.repo.Count()
	if err != nil {
		return fmt.Errorf("count failed: %w", err)
	}

	if count == 0 {
		return UserErrNoLists
	}

	lists, err := s.repo.GetAll()
	if err != nil {
		return fmt.Errorf("failed to get lists: %w", err)
	}

	clearScreen()
	printListHeader("lists", count)

	for _, l := range lists {
		printContent(l.Name, l.Id, count, formatter.Default, false)
	}

	return nil
}

func (s *Service) Add(name string) error {
	id, err := s.repo.Add(name)
	if err != nil {
		return fmt.Errorf("failed to add list: %w", err)
	}
	_, err = s.repo.GetActive()
	if err != nil {
		if errors.Is(err, ErrNoActiveList) {
			_, err := s.repo.SetActive(id, name)
			if err != nil {
				return fmt.Errorf("failed to set active list: %w", err)
			}
		} else {
			return fmt.Errorf("failed to get active list: %w", err)
		}
	}

	err = s.GetAll()
	if err != nil {
		return fmt.Errorf("failed to get lists: %w", err)
	}

	printColored("\nList successfully added.", formatter.Yellow)
	return nil
}

func (s *Service) Delete(id int) error {
	if err := s.repo.Delete(id); err != nil {
		if errors.Is(err, ErrNoRowsAffected) {
			return UserErrNoLists
		}
		return fmt.Errorf("failed to delete list: %w", err)
	}

	_, err := s.repo.GetActive()
	if err != nil {
		if errors.Is(err, ErrNoActiveList) {
			first, err := s.repo.GetFirst()
			if err != nil {
				if errors.Is(err, ErrNoRowsAffected) {
					printColored("\nList successfully deleted. No lists remain.", formatter.Yellow)
					return nil
				}
				return fmt.Errorf("failed to get first list: %w", err)
			}

			if _, err := s.repo.SetActive(first.Id, first.Name); err != nil {
				return fmt.Errorf("failed to set first list as active: %w", err)
			}
		} else {
			return fmt.Errorf("failed to get active list: %w", err)
		}
	}

	if err := s.GetAll(); err != nil {
		return fmt.Errorf("failed to get lists: %w", err)
	}

	printColored("\nList successfully deleted.", formatter.Yellow)
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
