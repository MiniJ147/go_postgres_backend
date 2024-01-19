-- +goose Up

CREATE TABLE authors (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name TEXT NOT NULL
);

CREATE TABLE books (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    title TEXT NOT NULL,
    author_id UUID,
    CONSTRAINT fk_author
        FOREIGN KEY(author_id)
            REFERENCES authors(id)
);

-- +goose Down

DROP TABLE authors;