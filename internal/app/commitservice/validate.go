package commitservice

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"log"
	"math/big"
)

// ValidateCommit выполняет базовую проверку валидности коммита:
// - проверка подписи
// - вычисление хэша
func (s *Service) ValidateCommit(input *CommitInput) (string, error) {
	log.Println("[validate] verifying signature and hashing commit")

	// 1. Собираем canonical commit body (без подписи)
	toHash := map[string]interface{}{
		"parent_hash":   input.ParentHash,
		"tree_hash":     input.TreeHash,
		"timestamp":     input.Timestamp,
		"author_pubkey": input.AuthorPubKey,
		"message":       input.Message,
	}
	jsonBytes, err := json.Marshal(toHash)
	if err != nil {
		return "", err
	}

	// 2. Вычисляем хэш
	hash := sha256.Sum256(jsonBytes)
	hashHex := hex.EncodeToString(hash[:])

	// 3. Проверяем подпись (ECDSA)
	valid, err := verifySignature(input.AuthorPubKey, input.Signature, hash[:])
	if err != nil {
		return "", err
	}
	if !valid {
		return "", errors.New("invalid_signature")
	}

	return hashHex, nil
}

func verifySignature(pubKeyHex string, sigHex string, msgHash []byte) (bool, error) {
	pubBytes, err := hex.DecodeString(pubKeyHex)
	if err != nil || len(pubBytes) != 65 || pubBytes[0] != 0x04 {
		return false, errors.New("invalid public key")
	}

	x := new(big.Int).SetBytes(pubBytes[1:33])
	y := new(big.Int).SetBytes(pubBytes[33:])
	pubKey := ecdsa.PublicKey{Curve: elliptic.P256(), X: x, Y: y}

	sigBytes, err := hex.DecodeString(sigHex)
	if err != nil || len(sigBytes) != 64 {
		return false, errors.New("invalid signature")
	}
	r := new(big.Int).SetBytes(sigBytes[:32])
	s := new(big.Int).SetBytes(sigBytes[32:])

	return ecdsa.Verify(&pubKey, msgHash, r, s), nil
}
