create table tasks
(
    task_id        integer default nextval('tasks_id_seq'::regclass) not null
        primary key,
    task_title     text,
    description    text,
    project_id_t   integer
        constraint tasks_project_id_fk
            references project
            on delete set null,
    start_time     timestamp with time zone,
    end_time       timestamp with time zone,
    type           integer,
    is_done        boolean default false,
    task_author_id integer
        constraint tasks_users_user_id_fk
            references users,
    is_public      boolean default false
);

alter table tasks
    owner to admin;

