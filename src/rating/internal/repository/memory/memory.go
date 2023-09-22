package memory

import (
	"context"
	"errors"

	"github.com/mrojasb2000/go-movie-app/src/rating/internal/repository"
	"github.com/mrojasb2000/go-movie-app/src/rating/pkg/model"
)

// ErrNotFound is returned when no ratings are found for a record.
var ErrNotFound = errors.New("ratings not found for a record")

// Repository defines a rating repository.
type Repository struct {
	data map[model.RecordType]map[model.RecordID][]model.Rating
}

// New creates a new memory repository.
func New() *Repository {
	return &Repository{map[model.RecordType]map[model.RecordID][]model.Rating{}}
}

// Get retrieves all rating for a given record.
func (r *Repository) Get(ctx context.Context, recordID model.RecordID, recordType model.RecordType) ([]model.Rating, error) {
	if _, ok := r.data[recordType]; !ok {
		return nil, repository.ErrNotFound
	}
	return r.data[recordType][recordID], nil
}

// Put adds a rating for a given record
func (r *Repository) Put(ctx context.Context, recordID model.RecordID, recordType model.RecordType, rating *model.Rating) error {
	if _, ok := r.data[recordType]; !ok {
		r.data[recordType] = map[model.RecordID][]model.Rating{}
	}
	r.data[recordType][recordID] = append(r.data[recordType][recordID], *rating)
	return nil
}
