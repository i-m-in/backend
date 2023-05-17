create table tasks_users
(
    id          serial
        primary key,
    task_id_ttu integer
        constraint tasks_users_tasks_task_id_fk
            references tasks,
    user_id_ttu integer
        constraint tasks_users_users_user_id_fk
            references users,
    role        text
);

alter table tasks_users
    owner to admin;

