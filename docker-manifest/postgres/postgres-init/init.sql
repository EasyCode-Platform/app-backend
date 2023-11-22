-- init app_backend
create database app_backend;

\c app_backend;

create user app_backend with encrypted password 'scut2023';

grant all privileges on database app_backend to app_backend;

CREATE EXTENSION pg_trgm;

CREATE EXTENSION btree_gin;