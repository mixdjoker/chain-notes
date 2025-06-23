package commitservice

import (
	"encoding/json"
	"log"

	"github.com/nats-io/nats.go"
)

func (s *Service) handleSubmit(msg *nats.Msg) {
	var input CommitInput
	if err := json.Unmarshal(msg.Data, &input); err != nil {
		s.respondReject(msg, "invalid_format", err.Error())
		return
	}

	log.Printf("[commit-service] received commit from %s", input.AuthorPubKey)

	hash, err := s.ValidateCommit(&input)
	if err != nil {
		s.respondReject(msg, "validation_failed", err.Error())
		return
	}

	// TODO: write to DB if valid

	// response success
	ack := CommitAccepted{
		Hash:         hash,
		Timestamp:    input.Timestamp,
		AuthorPubKey: input.AuthorPubKey,
	}
	data, _ := json.Marshal(ack)
	_ = msg.Respond(data)
}

func (s *Service) respondReject(msg *nats.Msg, reason, details string) {
	rej := CommitRejected{
		Error:   reason,
		Details: details,
	}
	data, _ := json.Marshal(rej)
	_ = msg.Respond(data)
}
