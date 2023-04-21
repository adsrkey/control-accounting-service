-- 000003_init_db.down.sql

drop table if exists development.operators cascade;
drop table if exists development.projects cascade;
drop table if exists development.project_operators cascade;

drop extension if exists pgcrypto;

drop type if exists project_type;

drop schema if exists development;
