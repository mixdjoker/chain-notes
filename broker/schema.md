# Основные события и темы в NATS

## `commit.submit` - Команда для отправки коммита в NATS.

Клиент отправляет зашифрованный коммит с подписью и метаданными.

```json
{
  "commit": {
    "parent_hash": "abc...",
    "tree_hash": "def...",
    "timestamp": "2025-06-21T...",
    "author_pubkey": "hex...",
    "signature": "sig...",
    "message": "Add new note"
  }
}
```

> Обрабатывается `CommitService`. Проверяется подпись, валидность parent_hash, хэш. Если всё корректно — сохраняется в базу

## `commit.accepted` - Сервис подтверждает приём коммита (всем подписчикам, например CLI, UI).

```json
{
  "hash": "ghi...",
  "timestamp": "...",
  "author_pubkey": "hex..."
}
```

> Отправляется `CommitService` после успешного сохранения коммита. Клиенты могут обновить своё состояние.

## `commit.rejected` - Сервис отклоняет коммит (например, из-за ошибки валидации).

```json
{
  "hash": "jkl...",
  "error": "Invalid signature",
  "details": "ECDSA verification failed"
}
```

> Отправляется `CommitService` при ошибке валидации. Клиенты могут уведомить пользователя об ошибке.

## `blob.store` - Клиент (или `CommitService`) загружает зашифрованный `blob`.

```json
{
  "hash": "abc123...",
  "type": "note",
  "mime_type": "text/plain",
  "data": "base64..."
}
```

> Можно сохранить в БД или перенаправить в IPFS.

## `tree.store` - Сервис сохраняет дерево из объектов.

```json
{
  "tree_hash": "def456...",
  "entries": [
    {
      "name": "note1.txt",
      "hash": "ghi789...",
      "type": "blob"
    },
    {
      "name": "note2.txt",
      "hash": "jkl012...",
      "type": "blob"
    }
  ]
}
```

## `chain.get_history` -> `chain.history` - запрос/ответ истории

Запрос на получение всей цепочки по `author_pubkey` (или последнего блока по ветке).

Запрос:
```json
{
  "author_pubkey": "hex...",
  "limit": 100
}
```

Ответ:
```json
{
  "author_pubkey": "hex...",
  "commits": [
    {
      "hash": "...",
      "parent_hash": "...",
      "tree_hash": "...",
      "timestamp": "...",
      "signature": "...",
      "message": "..."
    },
    ...
  ]
}
```

## `ref.update`

Обновление "указателя" на последний коммит ветки (аналог main, draft, tag)

```json
{
  "name": "main",
  "commit_hash": "abc..."
}
```

## Список тем по иерархии NATS


```
chain.commit.submit    -- Команда для отправки коммита в NATS.
chain.commit.accepted  -- Сервис подтверждает приём коммита (всем подписчикам)
chain.commit.rejected  -- Сервис отклоняет коммит (всем подписчикам)
chain.blob.store       -- Клиент/CommitService загружает зашифрованный blob
chain.tree.store       -- Сервис сохраняет дерево из объектов
chain.chain.get_history
chain.chain.history
chain.ref.update       -- Обновление "указателя" на последний коммит
```

Расширенный под пользователей:
```
user.{pubkey}.commit.accepted
user.{pubkey}.chain.history
```

## ⚠️ Безопасность

- Все события подписываются клиентом.
- Только валидированные сообщения обрабатываются.
- Можно использовать NATS JWT и подписи тем (subject auth) для доступа.