-- 000001_init_db.up.sql

create schema if not exists development;

create extension pgcrypto;

create table if not exists development.operators
(
    id                  uuid default gen_random_uuid() not null,
    created_at          timestamp                      not null,
    modified_at         timestamp                      not null,
    first_name          varchar(55)                    not null,
    last_name           varchar(55)                    not null,
    middle_name         varchar(55)                    not null,
    city                varchar(55)                    not null,
    country_code_number varchar(6)                     not null,
    phone_number        text                           not null,
    email               varchar(55) unique             not null,
    password            text                           not null
);

create type public.project_type as enum ('in', 'out', 'auto');

create table if not exists development.projects
(
    id           uuid default gen_random_uuid() not null,
    created_at   timestamp                      not null,
    modified_at  timestamp                      not null,
    project_name varchar(100) unique            not null,
    project_type project_type                   not null
);

alter table development.operators
    add constraint phone_number_chk
        check (phone_number ~ '^\d{10}$');

alter table development.projects
    add primary key (id);

alter table development.operators
    add primary key (id);

create table development.project_operators
(
    project_id  uuid references development.projects (id) on update cascade on delete cascade ,
    operator_id uuid references development.operators (id) on update cascade on delete cascade ,
    constraint project_operators_pkey primary key (project_id, operator_id)
);

create index on development.projects (id);
create index on development.operators (id);

create unique index uq_country_code_phone_number_operators
    on development.operators (country_code_number, phone_number);

create index on development.project_operators (project_id, operator_id);

create unique index uq_project_operators_ids
    on development.project_operators (project_id, operator_id);