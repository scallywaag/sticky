package notes

import "github.com/highseas-software/sticky/internal/lists"

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
