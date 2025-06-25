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

// Query represents a SQL query with its name and raw SQL string
type Query struct {
	Name     string
	QueryRaw string
}

func New(db *sql.DB) *Store {
	return &Store{db: db}
}

// ParentExists checks if a parent commit with the given hash exists
func (s *Store) ParentExists(ctx context.Context, hash string) (bool, error) {
	if hash == "" {
		return true, nil // root commit
	}

	q := Query{
		Name: "sqlstore.ParentExists",
	}

	query, args, err := sq.
		Select("1").
		From(tableCommits).
		Where(sq.Eq{colHash: hash}).
		Limit(1).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return false, fmt.Errorf("[sqlstore] failed to build SQL query %s: %w", q.Name, err)
	}

	q.QueryRaw = query

	var dummy int
	err = s.db.QueryRowContext(ctx, q.QueryRaw, args...).Scan(&dummy)
	if errors.Is(err, sql.ErrNoRows) {
		return false, nil // parent does not exist
	} else if err != nil {
		return false, fmt.Errorf("[sqlstore] failed to execute query %s: %w", q.Name, err)
	}

	return true, nil
}

// InsertCommit inserts a new commit into the database
func (s *Store) InsertCommit(ctx context.Context, hash string, input *commitservice.CommitInput) error {
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
