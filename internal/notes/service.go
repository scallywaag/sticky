package notes

import (
	"errors"
	"fmt"

	"github.com/highseas-software/sticky/internal/formatter"
	"github.com/highseas-software/sticky/internal/lists"
)

type Service struct {
	repo      Repository
	listsRepo lists.ActiveListRepo
}

func NewService(repo Repository, listsRepo lists.ActiveListRepo) *Service {
	return &Service{
		repo:      repo,
		listsRepo: listsRepo,
	}
}

type NotesResult struct {
	NotesList      []Note
	NotesCount     int
	ActiveListName string
}

func (s *Service) GetAll(listName string) (*NotesResult, error) {
	var activeList *lists.List
	var err error

	if listName == "" {
		activeList, err = s.listsRepo.GetActive()
		if err != nil {
			if errors.Is(err, lists.ErrNoActiveList) {
				return nil, lists.UserErrNoLists
			}
			return nil, fmt.Errorf("couldn't retrieve active list: %w", err)
		}
	} else {
		listId, err := s.listsRepo.GetId(listName)
		if err != nil {
			if errors.Is(err, lists.ErrListInexistent) {
				return nil, lists.UserErrInexistentList
			}
			return nil, fmt.Errorf("couldn't retrieve current list's id: %w", err)
		}

		err = s.listsRepo.SetActive(listId)
		if err != nil {
			return nil, fmt.Errorf("failed to set list as active: %w", err)
		}

		activeList = &lists.List{Id: listId, Name: listName}
	}

	count, err := s.repo.Count(activeList.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to count notes in active list: %w", err)
	}

	if count == 0 {
		return nil, UserErrNoNotes
	}

	notes, err := s.repo.GetAll(activeList.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to get notes: %w", err)
	}

	res := &NotesResult{
		NotesList:      notes,
		NotesCount:     count,
		ActiveListName: activeList.Name,
	}
	return res, nil
}

func (s *Service) Add(content string, color formatter.Color, status NoteStatus) (string, error) {
	activeList, err := s.listsRepo.GetActive()
	if err != nil {
		if errors.Is(err, lists.ErrNoActiveList) {
			return "", lists.UserErrNoLists
		}
		return "", fmt.Errorf("couldn't retrieve active list: %w", err)
	}

	n := NewNote(content, color, status)
	err = s.repo.Add(n, activeList.Id)
	if err != nil {
		return "", fmt.Errorf("failed to add note: %w", err)
	}

	return activeList.Name, nil
}

func (s *Service) Delete(id int) (string, error) {
	activeList, err := s.listsRepo.GetActive()
	if err != nil {
		return "", fmt.Errorf("failed to get active list: %w", err)
	}

	if exists, err := s.repo.CheckNotesExist(activeList.Id); err != nil {
		return "", fmt.Errorf("failed to check notes existence: %w", err)
	} else if !exists {
		return "", UserErrNoNotes
	}

	if noteExists, err := s.repo.CheckNoteExists(id, activeList.Id); err != nil {
		return "", fmt.Errorf("failed to check note existence: %w", err)
	} else if !noteExists {
		return "", UserErrInvalidDel
	}

	err = s.repo.Delete(id, activeList.Id)
	if err != nil {
		return "", fmt.Errorf("failed to delete note: %w", err)
	}

	return activeList.Name, nil
}

func (s *Service) Update(id int, color formatter.Color, status NoteStatus) (string, error) {
	activeList, err := s.listsRepo.GetActive()
	if err != nil {
		return "", fmt.Errorf("failed to get active list: %w", err)
	}

	if exists, err := s.repo.CheckNotesExist(activeList.Id); err != nil {
		return "", fmt.Errorf("failed to check note existence: %w", err)
	} else if !exists {
		return "", UserErrNoNotes
	}

	currentColor, currentStatus, err := s.repo.GetMutations(id, activeList.Id)
	if err != nil {
		if errors.Is(err, ErrNoRowsResult) {
			return "", UserErrInvalidMut
		}
		return "", fmt.Errorf("failed to get existing mutations: %w", err)
	}

	n := &Note{
		Id:     id,
		Color:  mutateColor(currentColor, color),
		Status: toggleStatus(currentStatus, status),
	}
	err = s.repo.Update(n, activeList.Id)
	if err != nil {
		return "", fmt.Errorf("failed to update note: %w", err)
	}

	return activeList.Name, nil
}
