package commitservice

import (
	"context"
)

// Store предоставляет интерфейс доступа к БД
type Store interface {
	ParentExists(ctx context.Context, hash string) (bool, error)
	InsertCommit(ctx context.Context, hash string, input *CommitInput) error
}
