CREATE TABLE commits (                   -- основа блокчейна
    hash STRING PRIMARY KEY,             -- SHA256(commit data)
    parent_hash STRING,                  -- NULL для корневого коммита
    tree_hash STRING NOT NULL,           -- корень содержимого
    timestamp TIMESTAMPTZ NOT NULL,
    author_pubkey STRING NOT NULL,
    signature STRING NOT NULL,
    message STRING,

    CONSTRAINT fk_parent FOREIGN KEY (parent_hash) REFERENCES commits(hash) ON DELETE SET NULL
);
