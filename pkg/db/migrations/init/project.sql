create table project
(
    project_id    integer default nextval('project_id_seq'::regclass) not null
        primary key,
    project_title text,
    maintainer_id integer
        constraint project_users_id_fk
            references users (id)
            on delete set null,
    color         text,
    description   text
);

alter table project
    owner to admin;

