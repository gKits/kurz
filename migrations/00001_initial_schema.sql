-- +goose Up
-- +goose StatementBegin
create table if not exists links (
    id integer primary key autoincrement,
    ref text not null,
    target text not null,
    created_at timestamp not null,
    created_by text not null,
    expires_at timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists links;
-- +goose StatementEnd
