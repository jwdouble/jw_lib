create table sys_template
(
    id          VARCHAR(128) primary key default gen_random_uuid(),
    name        varchar(255) not null default '',
    description varchar(255) not null default '',
    content     text         not null default '',
    delete_at   int not null default 0
);


-- postgres template must create manual
insert into sys_template
( name, content)
VALUES
('config-postgres-addr', 'host=${host} user=${user} password=${password} dbname=${dbname} port=${port} ')