// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: server_port.sql

package repo

import (
	"context"
)

const createServerPort = `-- name: CreateServerPort :one
INSERT INTO server_ports 
    (server_id, port)
VALUES
    (?, ?)
RETURNING server_id, port
`

type CreateServerPortParams struct {
	ServerID string `json:"server_id"`
	Port     int64  `json:"port"`
}

func (q *Queries) CreateServerPort(ctx context.Context, arg CreateServerPortParams) (ServerPort, error) {
	row := q.queryRow(ctx, q.createServerPortStmt, createServerPort, arg.ServerID, arg.Port)
	var i ServerPort
	err := row.Scan(&i.ServerID, &i.Port)
	return i, err
}

const getServerPort = `-- name: GetServerPort :one
SELECT server_id, port FROM server_ports WHERE server_id = ? AND port = ?
`

type GetServerPortParams struct {
	ServerID string `json:"server_id"`
	Port     int64  `json:"port"`
}

func (q *Queries) GetServerPort(ctx context.Context, arg GetServerPortParams) (ServerPort, error) {
	row := q.queryRow(ctx, q.getServerPortStmt, getServerPort, arg.ServerID, arg.Port)
	var i ServerPort
	err := row.Scan(&i.ServerID, &i.Port)
	return i, err
}

const listServerPorts = `-- name: ListServerPorts :many
SELECT server_id, port FROM server_ports
`

func (q *Queries) ListServerPorts(ctx context.Context) ([]ServerPort, error) {
	rows, err := q.query(ctx, q.listServerPortsStmt, listServerPorts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ServerPort{}
	for rows.Next() {
		var i ServerPort
		if err := rows.Scan(&i.ServerID, &i.Port); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
