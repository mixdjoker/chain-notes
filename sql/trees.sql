CREATE TABLE trees (
    hash STRING PRIMARY KEY,             -- SHA256(tree data)
    entry_name STRING NOT NULL,          -- "note.md"
    entry_type STRING NOT NULL,          -- "blob", "tree", "commit"
    entry_hash STRING NOT NULL           -- SHA256 от целевого объекта
);
