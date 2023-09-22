package memory

import (
	"context"
	"sync"

	"github.com/mrojasb2000/go-movie-app/src/metadata/internal/repository"
	"github.com/mrojasb2000/go-movie-app/src/metadata/pkg/model"
)

// Repository defines a memory movie metadata repository
type Repository struct {
	sync.RWMutex
	data map[string]*model.Metadata
}

// New creates a new memory repository
func New() *Repository {
	return &Repository{
		data: map[string]*model.Metadata{},
	}
}

// Get retrieves movie metadata for by movie id.
func (r *Repository) Get(_ context.Context, id string) (*model.Metadata, error) {
	r.RLock()
	defer r.RUnlock()
	m, ok := r.data[id]
	if !ok {
		return nil, repository.ErrNotFound
	}
	return m, nil
}

// Put adds movie metadat for a given movie id.
func (r *Repository) Put(_ context.Context, id string, metadata *model.Metadata) error {
	r.Lock()
	defer r.Unlock()
	r.data[id] = metadata
	return nil
}
