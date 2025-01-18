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