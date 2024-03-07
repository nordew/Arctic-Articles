-- +goose Up
-- +goose StatementBegin
CREATE TABLE articles (
    id UUID NOT NULL UNIQUE,
    title VARCHAR(70),
    content TEXT,
    author VARCHAR(32),
    image_url VARCHAR(255),
    date_published DATE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE articles IF EXISTS;
-- +goose StatementEnd
