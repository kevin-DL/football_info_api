-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
create table profiles
(
    id           uuid        not null
        constraint profiles_pk
            primary key,
    display_name varchar     not null,
    user_id uuid not null ,
    created_at   timestamptz not null,
    updated_at   timestamptz not null
);

create unique index profiles_user_id_index
    on profiles (user_id);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
drop table profiles;
-- +goose StatementEnd
