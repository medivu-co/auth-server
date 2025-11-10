create table public.clients
(
    id                    uuid                not null
        constraint client_pk_id
            primary key,
    name                  varchar             not null
        constraint client_ck_name
            unique,
    allowed_redirect_uris character varying[] not null
);