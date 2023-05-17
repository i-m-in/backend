create sequence users_id_seq
    as integer;

alter sequence users_id_seq owner to admin;

alter sequence users_id_seq owned by users.user_id;

