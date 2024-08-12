-- +goose Up
-- +goose StatementBegin
CREATE TABLE persons
(
    id serial,
    created_at timestamp(0) default NULL::timestamp without time zone,
    surname    varchar(100) not null,
    name       varchar(100) not null,
    gender     smallint     not null,
    birthday   date         not null,
    PRIMARY KEY (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table persons;
-- +goose StatementEnd