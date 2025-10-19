package commitservice

import "time"

// CommitInput represents an incoming message via NATS.
// Corresponds to the JSON message from chain.commit.submit
type CommitInput struct {
	ParentHash   string    `json:"parent_hash"`
	TreeHash     string    `json:"tree_hash"`
	Timestamp    time.Time `json:"timestamp"`
	AuthorPubKey string    `json:"author_pubkey"`
	Signature    string    `json:"signature"`
	Message      string    `json:"message"`
}

// CommitAccepted — message about a successful commit
type CommitAccepted struct {
	Hash         string    `json:"hash"`
	Timestamp    time.Time `json:"timestamp"`
	AuthorPubKey string    `json:"author_pubkey"`
}

// CommitRejected — message about an error
type CommitRejected struct {
	Hash    string `json:"hash"`
	Error   string `json:"error"`
	Details string `json:"details,omitempty"`
}
