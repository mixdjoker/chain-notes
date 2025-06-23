CREATE TABLE commit_refs (
    name STRING PRIMARY KEY,            -- main, draft, test
    commit_hash STRING NOT NULL,        -- актуальный хэш
    updated_at TIMESTAMPTZ
);
