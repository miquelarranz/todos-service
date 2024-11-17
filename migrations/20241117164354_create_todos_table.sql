-- +goose Up
CREATE TABLE todos (
    id text NOT NULL,
    title text,
    description text,
    PRIMARY KEY(id)
);

-- +goose Down
DROP TABLE todos;