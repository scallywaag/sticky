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

func (s *Service) GetAll() ([]List, int, error) {
	count, err := s.repo.Count()
	if err != nil {
		return nil, 0, fmt.Errorf("count failed: %w", err)
	}

	if count == 0 {
		return nil, 0, UserErrNoLists
	}

	lists, err := s.repo.GetAll()
	if err != nil {
		return nil, 0, fmt.Errorf("failed to get lists: %w", err)
	}

	return lists, count, nil
}

func (s *Service) Add(name string) error {
	id, err := s.repo.Add(name)
	if err != nil {
		return fmt.Errorf("failed to add list: %w", err)
	}

	err = s.repo.SetActive(id)
	if err != nil {
		return fmt.Errorf("failed to set active list: %w", err)
	}

	return nil
}

func (s *Service) Delete(id int) error {
	if err := s.repo.Delete(id); err != nil {
		if errors.Is(err, ErrNoRowsAffected) {
			return UserErrNoLists
		}
		return fmt.Errorf("failed to delete list: %w", err)
	}

	_, err := s.GetActiveOrSetFirst()
	if err != nil {
		return fmt.Errorf("failed to get or set active list: %w", err)
	}

	return nil
}

func (s *Service) GetActiveOrSetFirst() (string, error) {
	_, err := s.repo.GetActive()
	if err != nil {
		if errors.Is(err, ErrNoActiveList) {
			first, err := s.repo.GetFirst()
			if err != nil {
				if errors.Is(err, ErrNoListsExist) {
					return "", ErrNoListsExist
				}
				return "", fmt.Errorf("failed to get first list: %w", err)
			}

			if err := s.repo.SetActive(first.Id); err != nil {
				return "", fmt.Errorf("failed to set first list as active: %w", err)
			}
		} else {
			return "", fmt.Errorf("failed to get active list: %w", err)
		}
	}

	return "", nil
}
