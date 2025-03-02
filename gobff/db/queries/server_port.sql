-- name: CreateServerPort :one
INSERT INTO server_ports 
    (server_id, port)
VALUES
    (?, ?)
RETURNING *;

-- name: GetServerPort :one
SELECT * FROM server_ports WHERE server_id = ? AND port = ?;

-- name: ListServerPorts :many
SELECT * FROM server_ports;

-- name: CreateServerGroup :one
INSERT INTO server_groups 
    (name, desc)
VALUES
    (?, ?)
RETURNING *;

-- name: GetServerGroup :one
SELECT id, name, desc FROM server_groups WHERE id = ?;

-- name: ListServerGroups :many
SELECT id, name, desc FROM server_groups;

-- name: DeleteServerGroup :exec
DELETE FROM server_groups WHERE name = ?;

-- name: CreateServer :one
INSERT INTO servers 
    (id, name, ip, group_id, reg_status)
VALUES
    (?, ?, ?, ?, 'NEW')
RETURNING id;

-- name: ListServersByGroup :many
SELECT * FROM servers WHERE group_id = ?;

-- name: GetServerByGidSid :one
SELECT * FROM servers WHERE group_id = ? AND id = ?;

-- name: UpdateOneTimeTokenForServerRegistration :exec
UPDATE servers SET one_time_token = ?, reg_status = 'PENDING' WHERE id = ?;

-- name: GetOneTimeTokenForServerRegistration :one
SELECT one_time_token FROM servers WHERE id = ?;
