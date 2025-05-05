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
    id text not null, -- group_id::ip
    group_id INTEGER not null,
    ip text not null,
    mac text, -- UUID from server cat /etc/machine-id 
    reg_status text not null, -- NEW -> PENDING -> ACTIVE
    one_time_token text,
    one_time_token_expiry INTEGER,
    name text not null,
    desc text,
    fqdn text,
    agent_version text,
    os text, -- OS name and version
    arch text, -- system arch viz. x86_64, arm64
    nats_subject text, -- NATs subject to which remote agent will publish stats
    monit_enabled BOOLEAN, -- if monit is enabled on the remote host; overrides group setting
    notifs_enabled BOOLEAN, -- if notifications are enabled on the remote host; overrides group setting
    FOREIGN KEY(group_id) REFERENCES server_groups(id),
    PRIMARY KEY(group_id, ip)
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

create table if not exists http_monit_configs(
    encoded_url text PRIMARY KEY,
    url text unique not null,
    friendly_name text unique not null,
    interval INTEGER not null, -- check every interval seconds
    retries INTEGER not null, -- Maximum retries before the service is marked as down and a notification is sent
    timeout INTEGER not null,
    upside_down BOOLEAN not null, -- if true, the service is considered down if the response is not equal to expected_response
    max_redirects INTEGER not null,
    method text not null,
    accepted_codes text not null,
    body_encoding text not null,
    body text,
    headers text,
    authentication_mode text,
    expected_response text not null
);

create table http_monit_data(
    url text not null,
    timestamp INTEGER not null,
    up BOOLEAN not null,
    status_code INTEGER,
    response_time INTEGER,
    response_body text,
    FOREIGN KEY(url) REFERENCES http_monit_configs(url),
    primary key(url, timestamp) 
);