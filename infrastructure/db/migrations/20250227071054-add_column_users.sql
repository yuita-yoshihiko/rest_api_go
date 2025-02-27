
-- +migrate Up
alter table users add column email text;
-- +migrate Down
alter table users drop column email;
