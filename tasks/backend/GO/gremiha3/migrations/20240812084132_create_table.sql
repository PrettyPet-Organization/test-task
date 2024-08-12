-- +goose Up
-- +goose StatementBegin
SELECT setval(pg_get_serial_sequence('persons', 'id'), max(id)) FROM persons;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
