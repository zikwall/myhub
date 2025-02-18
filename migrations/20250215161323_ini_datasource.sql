-- +goose Up
-- +goose StatementBegin
create table channels
(
    pid         serial,
    id          uuid         not null,
    name        varchar(255) not null,
    description text         null,
    created_at  timestamp default now(),
    updated_at  timestamp default now(),
    deleted_at  timestamp    null,
    primary key (id)
);

create table timezones
(
    pid        serial,
    id         uuid      not null,
    zone       smallint,
    zone_rfc   varchar(30),
    name       varchar(10),
    created_at timestamp default now(),
    updated_at timestamp default now(),
    deleted_at timestamp null,
    primary key (id)
);

create table categories
(
    pid        serial,
    id         uuid      not null,
    name       varchar(255),
    created_at timestamp default now(),
    updated_at timestamp default now(),
    deleted_at timestamp null,
    primary key (id)
);

create table countries
(
    pid        serial,
    id         uuid         not null,
    name       varchar(255) not null,
    iso3366_1  varchar(3)   not null,
    created_at timestamp default now(),
    updated_at timestamp default now(),
    deleted_at timestamp    null,
    primary key (id)
);

create table tags
(
    pid        serial,
    id         uuid        not null,
    label      varchar(50) not null,
    created_at timestamp default now(),
    updated_at timestamp default now(),
    deleted_at timestamp   null,
    primary key (id)
);

create table users
(
    pid           serial,
    id            uuid         not null,
    user_name     varchar(50)  not null,
    first_name    varchar(100) not null,
    last_name     varchar(100) null,
    email         varchar      not null,
    avatar_b64    text         not null default '',
    avatar_url    text         not null default '',
    password_hash varchar(255) not null,
    created_at    timestamp             default now(),
    updated_at    timestamp             default now(),
    deleted_at    timestamp    null,
    primary key (id)
);

create table channel_user
(
    pid         serial,
    id          uuid         not null,
    channel_id  uuid         not null,
    zone_id     uuid         not null,
    category_id uuid         not null,
    user_id     uuid         not null,
    name        varchar(255) not null,
    url         text         not null,
    countries   uuid[]       null,
    tags        uuid[]       null,
    created_at  timestamp default now(),
    updated_at  timestamp default now(),
    deleted_at  timestamp    null,
    primary key (id)
);

alter table channel_user
    add constraint fk_channel_user_channel
        foreign key (channel_id)
            references channels (id)
            on delete cascade;

alter table channel_user
    add constraint fk_channel_user_zone
        foreign key (zone_id)
            references timezones (id)
            on delete cascade;

alter table channel_user
    add constraint fk_channel_user_category
        foreign key (category_id)
            references countries (id)
            on delete cascade;

alter table channel_user
    add constraint fk_channel_user_user
        foreign key (user_id)
            references users (id)
            on delete cascade;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter table channel_user
    drop constraint fk_channel_user_channel;

alter table channel_user
    drop constraint fk_channel_user_zone;

alter table channel_user
    drop constraint fk_channel_user_category;

alter table channel_user
    drop constraint fk_channel_user_user;

drop table channel_user;
-- Удаление таблицы timezones
drop table timezones;
-- Удаление таблицы channel_category
drop table categories;
-- Удаление таблицы countries
drop table countries;
-- Удаление таблицы tags
drop table tags;
-- Удаление таблицы channel
drop table channels;
-- Удаление таблицы users
drop table users;
-- +goose StatementEnd
