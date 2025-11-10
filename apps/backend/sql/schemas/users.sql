create table public.users
(
    id            serial
        constraint user_pk_id
            primary key,
    email         varchar                                not null
        constraint users_ck_email
            unique,
    password_hash char(60)                               not null,
    name          varchar(25)                            not null,
    created_at    timestamp with time zone default now() not null
);