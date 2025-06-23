package commitservice

import "time"

// CommitInput — структура входящего сообщения через NATS
// Соответствует JSON-сообщению из chain.commit.submit

type CommitInput struct {
	ParentHash   string    `json:"parent_hash"`
	TreeHash     string    `json:"tree_hash"`
	Timestamp    time.Time `json:"timestamp"`
	AuthorPubKey string    `json:"author_pubkey"`
	Signature    string    `json:"signature"`
	Message      string    `json:"message"`
}

// CommitAccepted — сообщение об успешном коммите

type CommitAccepted struct {
	Hash         string    `json:"hash"`
	Timestamp    time.Time `json:"timestamp"`
	AuthorPubKey string    `json:"author_pubkey"`
}

// CommitRejected — сообщение об ошибке

type CommitRejected struct {
	Hash    string `json:"hash"`
	Error   string `json:"error"`
	Details string `json:"details,omitempty"`
}
