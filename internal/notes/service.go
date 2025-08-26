package notes

import (
	"fmt"

	"github.com/highseas-software/sticky/internal/formatter"
	"github.com/highseas-software/sticky/internal/lists"
)

type Service struct {
	repo      Repository
	listsRepo lists.Repository
}

func NewService(repo Repository, listsRepo lists.Repository) *Service {
	return &Service{
		repo:      repo,
		listsRepo: listsRepo,
	}
}

func (s *Service) GetAll(listName string) error {
	var activeList *lists.List
	var err error

	if listName == "" {
		activeList, err = s.listsRepo.GetActive()
		if err != nil {
			return fmt.Errorf("couldn't retrieve active list: %w", err)
		}
	} else {
		listId, err := s.listsRepo.GetId(listName)
		if err != nil {
			return fmt.Errorf("couldn't retrieve current list's id: %w", err)
		}

		activeList, err = s.listsRepo.SetActive(listId, listName)
		if err != nil {
			return fmt.Errorf("failed to set list as active: %w", err)
		}
	}

	count, err := s.repo.Count(activeList.Id)
	if err != nil {
		return fmt.Errorf("failed to count notes in active list: %w", err)
	}

	notes, err := s.repo.GetAll(activeList.Id)
	if err != nil {
		return fmt.Errorf("failed to get notes: %w", err)
	}

	formatter.ClearScreen()
	formatter.PrintListHeader(activeList.Name, count)

	for _, n := range notes {
		cross := n.Status == StatusCross
		formatter.Print(n.Content, n.Id, count, n.Color, cross)
	}

	return nil
}

func (s *Service) Delete(id int) error {
	activeList, err := s.listsRepo.GetActive()
	if err != nil {
		return fmt.Errorf("failed to get active list: %w", err)
	}

	err = s.repo.Delete(id, activeList.Id)
	if err != nil {
		return fmt.Errorf("failed to get delete note: %w", err)
	}

	s.GetAll(activeList.Name)

	formatter.PrintColored("\nNote successfully deleted.", formatter.Yellow)
	return nil
}
