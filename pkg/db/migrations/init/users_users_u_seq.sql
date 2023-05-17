create sequence users_users_u_seq
    as integer;

alter sequence users_users_u_seq owner to admin;

alter sequence users_users_u_seq owned by users_users.id;

