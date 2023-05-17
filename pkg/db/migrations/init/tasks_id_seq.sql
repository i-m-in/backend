create sequence tasks_id_seq
    as integer;

alter sequence tasks_id_seq owner to admin;

alter sequence tasks_id_seq owned by tasks.task_id;

