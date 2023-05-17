create table users_users
(
    id       integer default nextval('users_users_u_seq'::regclass) not null
        primary key,
    user1_id integer,
    user2_id integer
        constraint users_users_users_user_id_fk
            references users,
    state    integer
);

alter table users_users
    owner to admin;

