CREATE TABLE blobs (
    hash STRING PRIMARY KEY,
    type STRING NOT NULL,                -- "note", "file", ...
    mime_type STRING NOT NULL,
    encrypted_data BYTES NOT NULL        -- base64(AES(...))
);
