-- Active: 1714057451806@@127.0.0.1@5432@restapi_dev

CREATE TABLE users (
    id bigserial not null primary key,
    email varchar not null unique,
    encrypted_password varchar not null
);