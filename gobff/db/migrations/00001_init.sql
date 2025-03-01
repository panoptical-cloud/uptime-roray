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

create table if not exists inflight_servers(
    id INTEGER PRIMARY KEY,
    name text not null,
    desc text not null,
    fqdn text not null UNIQUE,
    status text not null, -- NEW -> PENDING -> ACTIVE
    one_time_token text,
    one_time_token_expiry INTEGER,
    group_id INTEGER not null,
    FOREIGN KEY(group_id) REFERENCES server_groups(id)
)

create table if not exists servers (
    id text primary key, -- UUID from server cat /etc/machine-id
    inflight_servers_id INTEGER,
    group_id INTEGER not null,
    name text not null,
    desc text not null,
    fqdn text not null,
    ip text,
    agent_version text,
    os text, -- OS name and version
    arch text, -- system arch viz. x86_64, arm64
    nats_subject text, -- NATs subject to which remote agent will publish stats
    monit_enabled BOOLEAN, -- if monit is enabled on the remote host; overrides group setting
    notifs_enabled BOOLEAN, -- if notifications are enabled on the remote host; overrides group setting
    FOREIGN KEY(group_id) REFERENCES server_groups(id),
    FOREIGN KEY(inflight_servers_id) REFERENCES inflight_servers(id)
);

create table if not exists server_configs(
    server_id text not null,
    key text not null,
    value text not null,
    FOREIGN KEY(server_id) REFERENCES servers(id),
    primary key(server_id, key)
);

create table if not exists server_up ( -- central server/executor calls host to check if it is reachable
    server_id text not null,
    timestamp INTEGER not null,
    up BOOLEAN not null,
    uptime INTEGER,
    FOREIGN KEY(server_id) REFERENCES servers(id),
    primary key(server_id, timestamp)
);

create table if not exists server_base_stats(
    server_id text not null,
    timestamp INTEGER not null,
    cpu_usage REAL,
    memory_usage REAL,
    disk_usage REAL,
    network_usage REAL,
    FOREIGN KEY(server_id) REFERENCES servers(id),
    primary key(server_id, timestamp)
);

create table if not exists server_ports_up(
    server_id text not null,
    port INTEGER not null,
    timestamp INTEGER not null,
    up BOOLEAN not null,
    FOREIGN KEY(server_id) REFERENCES servers(id),
    primary key(server_id, port, timestamp)
);