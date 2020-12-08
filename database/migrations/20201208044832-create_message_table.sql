
-- +migrate Up
create table if not exists message
(
    id      int primary key,
    content text not null
);

-- +migrate Down
drop table if exists message;
