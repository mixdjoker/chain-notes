package domain

import "time"

type Blob struct {
	Type     string // "note", "file", "json", "markdown", ...
	MimeType string // например, "text/plain", "application/json"
	Data     []byte // Зашифрованные данные (AES-GCM, Base64)
}

type Tree struct {
	Entries []TreeEntry
}

type TreeEntry struct {
	Name string // имя файла / ключа
	Hash string // hash(blob or commit or tree)
	Type string // "blob" | "tree" | "commit"
}

type Commit struct {
	Hash         string  // SHA-256 от содержимого коммита
	ParentHash   *string // Предыдущий коммит (или nil для корня)
	TreeHash     string  // Корень содержимого (может быть одной заметкой или деревом)
	Timestamp    time.Time
	AuthorPubKey string // Публичный ключ автора
	Signature    string // Подпись: Sign(hash(tree + parent + ts))
	Message      string // (необязательное) описание коммита
}
