-- +goose Up
-- +goose StatementBegin
CREATE TABLE obligation (
    id SERIAL,
    title VARCHAR(255),
    description text,
    PRIMARY KEY(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE obligation;
-- +goose StatementEnd
