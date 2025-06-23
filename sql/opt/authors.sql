CREATE TABLE authors (              -- публичные ключи пользователей (для индексации)
    pubkey STRING PRIMARY KEY,
    name STRING,
    email STRING,
    registered_at TIMESTAMPTZ
);
