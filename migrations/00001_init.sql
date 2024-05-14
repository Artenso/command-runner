-- +goose Up
-- +goose StatementBegin

CREATE TYPE status AS ENUM ('UNSPECIFIED', 'NEW', 'IN_PROGRESS', 'DONE', 'FAILED', 'STOPPED');

CREATE TABLE commands (
    id BIGSERIAL PRIMARY KEY ,
    command TEXT NOT NULL DEFAULT '',
    status status NOT NULL DEFAULT 'UNSPECIFIED',
    pid BIGINT,
    output TEXT
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE commands;

-- +goose StatementEnd