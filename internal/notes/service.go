package notes

import (
	"errors"
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
	var activeListId *lists.List
	var err error

	if listName == "" {
		activeListId, err = s.listsRepo.GetActive()
		if err != nil {
			if errors.Is(err, lists.ErrNoActiveList) {
				return lists.UserErrNoLists
			}
			return fmt.Errorf("couldn't retrieve active list: %w", err)
		}
	} else {
		listId, err := s.listsRepo.GetId(listName)
		if err != nil {
			return fmt.Errorf("couldn't retrieve current list's id: %w", err)
		}

		activeListId, err = s.listsRepo.SetActive(listId, listName)
		if err != nil {
			return fmt.Errorf("failed to set list as active: %w", err)
		}
	}

	count, err := s.repo.Count(activeListId.Id)
	if err != nil {
		return fmt.Errorf("failed to count notes in active list: %w", err)
	}

	notes, err := s.repo.GetAll(activeListId.Id)
	if err != nil {
		return fmt.Errorf("failed to get notes: %w", err)
	}

	formatter.ClearScreen()
	formatter.PrintListHeader(activeListId.Name, count)

	for _, n := range notes {
		cross := n.Status == StatusCross
		formatter.PrintContent(n.Content, n.Id, count, n.Color, cross)
	}

	return nil
}

func (s *Service) Add(content string, color formatter.Color, status NoteStatus) error {
	activeList, err := s.listsRepo.GetActive()
	if err != nil {
		return fmt.Errorf("failed to get active list: %w", err)
	}

	c := color
	if c == "" {
		c = formatter.Default
	}

	x := status
	if x == "" {
		x = StatusDefault
	}
	n := &Note{Content: content, Color: c, Status: x}
	err = s.repo.Add(n, activeList.Id)
	if err != nil {
		return fmt.Errorf("failed to add note: %w", err)
	}

	err = s.GetAll(activeList.Name)
	if err != nil {
		return fmt.Errorf("could not retrieve notes list: %w", err)
	}

	formatter.PrintColored("\nNote successfully added.", formatter.Yellow)
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

	err = s.GetAll(activeList.Name)
	if err != nil {
		return fmt.Errorf("could not retrieve notes list: %w", err)
	}

	formatter.PrintColored("\nNote successfully deleted.", formatter.Yellow)
	return nil
}

func (s *Service) Update(id int, color formatter.Color, status NoteStatus) error {
	activeList, err := s.listsRepo.GetActive()
	if err != nil {
		return fmt.Errorf("failed to get active list: %w", err)
	}

	currentColor, currentStatus, err := s.repo.GetMutations(id, activeList.Id)
	if err != nil {
		return fmt.Errorf("failed to get existing mutations: %w", err)
	}

	c := mutateColor(currentColor, color)
	x := toggleStatus(currentStatus, status)

	n := &Note{Id: id, Color: c, Status: x}
	err = s.repo.Update(n, activeList.Id)
	if err != nil {
		return fmt.Errorf("failed to update note: %w", err)
	}

	err = s.GetAll(activeList.Name)
	if err != nil {
		return fmt.Errorf("could not retrieve notes list: %w", err)
	}

	formatter.PrintColored("\nNote successfully mutated.", formatter.Yellow)
	return nil
}
