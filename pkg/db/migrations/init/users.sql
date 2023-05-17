create table users
(
    user_id    integer default nextval('users_id_seq'::regclass) not null
        primary key,
    user_name  text,
    first_name text,
    last_name  text,
    email      text
        unique,
    password   text
);

alter table users
    owner to admin;

