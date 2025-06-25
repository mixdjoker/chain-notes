package sqlstore

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	sq "github.com/Masterminds/squirrel"

	"github.com/mixdjoker/chain-notes/internal/app/commitservice"
)

const (
	tableCommits    = "commits"
	colHash         = "hash"
	colParentHash   = "parent_hash"
	colTreeHash     = "tree_hash"
	colTimestamp    = "timestamp"
	colAuthorPubKey = "author_pubkey"
	colSignature    = "signature"
	colMessage      = "message"
)

// Store implements the commitservice.Store interface using a SQL database
type Store struct {
	db *sql.DB
}

func New(db *sql.DB) *Store {
	return &Store{db: db}
}

// ParentExists checks if a parent commit with the given hash exists
func (s *Store) ParentExists(ctx context.Context, hash string) (bool, error) {
	if hash == "" {
		return true, nil // root commit
	}

	queryName := "sqlstore.ParentExists"

	query, args, err := sq.
		Select("1").
		From(tableCommits).
		Where(sq.Eq{colHash: hash}).
		Limit(1).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return false, fmt.Errorf("[sqlstore] failed to build SQL query %s: %w", queryName, err)
	}

	var dummy int
	err = s.db.QueryRowContext(ctx, query, args...).Scan(&dummy)
	if errors.Is(err, sql.ErrNoRows) {
		return false, nil // parent does not exist
	} else if err != nil {
		return false, fmt.Errorf("[sqlstore] failed to execute query %s: %w", queryName, err)
	}

	return true, nil
}

// InsertCommit inserts a new commit into the database
func (s *Store) InsertCommit(ctx context.Context, hash string, input *commitservice.CommitInput) error {
	queryName := "sqlstore.InsertCommit"

	commitColums := []string{
		colHash,
		colParentHash,
		colTreeHash,
		colTimestamp,
		colAuthorPubKey,
		colSignature,
		colMessage,
	}

	query, args, err := sq.
		Insert(tableCommits).
		Columns(commitColums...).
		Values(
			hash,
			input.ParentHash,
			input.TreeHash,
			input.Timestamp,
			input.AuthorPubKey,
			input.Signature,
			input.Message,
		).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return fmt.Errorf("[sqlstore] failed to build SQL query %s: %w", queryName, err)
	}

	_, err = s.db.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("[sqlstore] failed to insert commit %s: %w", queryName, err)
	}

	return nil
}
