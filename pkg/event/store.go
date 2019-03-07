package event

import (
	"context"

	"github.com/hellofresh/goengine/aggregate"
)

// Store is a repository for aggregate root
type Store struct {
	repo *aggregate.Repository
}

// NewStore create a new EventStore instance
func NewStore(repo *aggregate.Repository) *Store {
	return &Store{
		repo: repo,
	}
}

// Save saves the aggregate root
func (s *Store) Save(ctx context.Context, root aggregate.Root) error {
	return s.repo.SaveAggregateRoot(ctx, root)
}

// Get loads the aggregate root
func (s *Store) Get(ctx context.Context, aggregateID aggregate.ID) (aggregate.Root, error) {
	return s.repo.GetAggregateRoot(ctx, aggregateID)
}
