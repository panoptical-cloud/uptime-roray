create table if not exists server_ports(
    server_id text not null,
    port SMALLINT not null,
    primary key(server_id, port)
);