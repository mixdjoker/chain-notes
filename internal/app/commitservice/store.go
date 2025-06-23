package commitservice

import (
	"context"
	"database/sql"
	"errors"
)

// Store предоставляет интерфейс доступа к БД
type Store interface {
	ParentExists(ctx context.Context, hash string) (bool, error)
	InsertCommit(ctx context.Context, hash string, input *CommitInput) error
}

// SQLStore — реализация Store через sql.DB (CockroachDB)
type SQLStore struct {
	db *sql.DB
}

// NewSQLStore создает новый экземпляр SQLStore с заданным sql.DB
func NewSQLStore(db *sql.DB) *SQLStore {
	return &SQLStore{db: db}
}

// ParentExists проверяет, существует ли родительский коммит с заданным хэшем
// Возвращает true, если коммит родительский (пустой хэш)
func (s *SQLStore) ParentExists(ctx context.Context, hash string) (bool, error) {
	if hash == "" {
		return true, nil // root commit
	}
	var exists bool
	err := s.db.QueryRowContext(ctx, `SELECT EXISTS (SELECT 1 FROM commits WHERE hash = $1)`, hash).Scan(&exists)
	return exists, err
}

// InsertCommit вставляет новый коммит в базу данных
// Возвращает ошибку, если вставка не удалась
func (s *SQLStore) InsertCommit(ctx context.Context, hash string, input *CommitInput) error {
	_, err := s.db.ExecContext(ctx, `
		INSERT INTO commits (hash, parent_hash, tree_hash, timestamp, author_pubkey, signature, message)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`,
		hash,
		input.ParentHash,
		input.TreeHash,
		input.Timestamp,
		input.AuthorPubKey,
		input.Signature,
		input.Message,
	)
	if err != nil {
		return errors.New("failed to insert commit: " + err.Error())
	}
	return nil
}
