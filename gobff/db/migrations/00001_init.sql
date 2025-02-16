create table if not exists server_ports(
    server_id text not null,
    port SMALLINT not null,
    primary key(server_id, port)
);

create table if not exists server_groups(
    id INTEGER PRIMARY KEY,
    name text unique not null,
    desc text not null
);

create table if not exists servers (
    id INTEGER primary key,
    name text not null,
    hostname text not null,
    ip text,
    agent_port SMALLINT,
    agent_version text,
    group_id INTEGER not null,
    one_time_token text,
    FOREIGN KEY(group_id) REFERENCES server_groups(id)
);