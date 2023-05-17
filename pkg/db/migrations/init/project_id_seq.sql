create sequence project_id_seq
    as integer;

alter sequence project_id_seq owner to admin;

alter sequence project_id_seq owned by project.project_id;

