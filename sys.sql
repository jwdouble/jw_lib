create table sys_template
(
    id          VARCHAR(128) primary key default gen_random_uuid(),
    name        varchar(255) not null,
    description varchar(255) not null,
    content     text         not null
);