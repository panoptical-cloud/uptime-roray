// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: server_port.sql

package repo

import (
	"context"
)

const completeServerRegistration = `-- name: CompleteServerRegistration :exec
UPDATE servers SET reg_status = 'ACTIVE', agent_version = ?, mac = ? WHERE id = ?
`

type CompleteServerRegistrationParams struct {
	AgentVersion *string `json:"agent_version"`
	Mac          *string `json:"mac"`
	ID           string  `json:"id"`
}

func (q *Queries) CompleteServerRegistration(ctx context.Context, arg CompleteServerRegistrationParams) error {
	_, err := q.exec(ctx, q.completeServerRegistrationStmt, completeServerRegistration, arg.AgentVersion, arg.Mac, arg.ID)
	return err
}

const createHttpUrlConfig = `-- name: CreateHttpUrlConfig :one
INSERT INTO http_monit_configs
    (url, friendly_name, interval, retries, timeout, upside_down, max_redirects, method, accepted_codes, body_encoding, body, headers, authentication_mode, expected_response)
VALUES
    (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
RETURNING encoded_url, url, friendly_name, interval, retries, timeout, upside_down, max_redirects, method, accepted_codes, body_encoding, body, headers, authentication_mode, expected_response
`

type CreateHttpUrlConfigParams struct {
	Url                string  `json:"url"`
	FriendlyName       string  `json:"friendly_name"`
	Interval           int64   `json:"interval"`
	Retries            int64   `json:"retries"`
	Timeout            int64   `json:"timeout"`
	UpsideDown         bool    `json:"upside_down"`
	MaxRedirects       int64   `json:"max_redirects"`
	Method             string  `json:"method"`
	AcceptedCodes      string  `json:"accepted_codes"`
	BodyEncoding       string  `json:"body_encoding"`
	Body               *string `json:"body"`
	Headers            *string `json:"headers"`
	AuthenticationMode *string `json:"authentication_mode"`
	ExpectedResponse   string  `json:"expected_response"`
}

func (q *Queries) CreateHttpUrlConfig(ctx context.Context, arg CreateHttpUrlConfigParams) (HttpMonitConfig, error) {
	row := q.queryRow(ctx, q.createHttpUrlConfigStmt, createHttpUrlConfig,
		arg.Url,
		arg.FriendlyName,
		arg.Interval,
		arg.Retries,
		arg.Timeout,
		arg.UpsideDown,
		arg.MaxRedirects,
		arg.Method,
		arg.AcceptedCodes,
		arg.BodyEncoding,
		arg.Body,
		arg.Headers,
		arg.AuthenticationMode,
		arg.ExpectedResponse,
	)
	var i HttpMonitConfig
	err := row.Scan(
		&i.EncodedUrl,
		&i.Url,
		&i.FriendlyName,
		&i.Interval,
		&i.Retries,
		&i.Timeout,
		&i.UpsideDown,
		&i.MaxRedirects,
		&i.Method,
		&i.AcceptedCodes,
		&i.BodyEncoding,
		&i.Body,
		&i.Headers,
		&i.AuthenticationMode,
		&i.ExpectedResponse,
	)
	return i, err
}

const createServer = `-- name: CreateServer :one
INSERT INTO servers 
    (id, name, ip, group_id, reg_status)
VALUES
    (?, ?, ?, ?, 'NEW')
RETURNING id
`

type CreateServerParams struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Ip      string `json:"ip"`
	GroupID int64  `json:"group_id"`
}

func (q *Queries) CreateServer(ctx context.Context, arg CreateServerParams) (string, error) {
	row := q.queryRow(ctx, q.createServerStmt, createServer,
		arg.ID,
		arg.Name,
		arg.Ip,
		arg.GroupID,
	)
	var id string
	err := row.Scan(&id)
	return id, err
}

const createServerGroup = `-- name: CreateServerGroup :one
INSERT INTO server_groups 
    (name, desc)
VALUES
    (?, ?)
RETURNING id, name, "desc"
`

type CreateServerGroupParams struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
}

func (q *Queries) CreateServerGroup(ctx context.Context, arg CreateServerGroupParams) (ServerGroup, error) {
	row := q.queryRow(ctx, q.createServerGroupStmt, createServerGroup, arg.Name, arg.Desc)
	var i ServerGroup
	err := row.Scan(&i.ID, &i.Name, &i.Desc)
	return i, err
}

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

const deleteServerGroup = `-- name: DeleteServerGroup :exec
DELETE FROM server_groups WHERE name = ?
`

func (q *Queries) DeleteServerGroup(ctx context.Context, name string) error {
	_, err := q.exec(ctx, q.deleteServerGroupStmt, deleteServerGroup, name)
	return err
}

const getHttpUrlConfigById = `-- name: GetHttpUrlConfigById :one
SELECT encoded_url, url, friendly_name, interval, retries, timeout, upside_down, max_redirects, method, accepted_codes, body_encoding, body, headers, authentication_mode, expected_response FROM http_monit_configs WHERE encoded_url = ?
`

func (q *Queries) GetHttpUrlConfigById(ctx context.Context, encodedUrl string) (HttpMonitConfig, error) {
	row := q.queryRow(ctx, q.getHttpUrlConfigByIdStmt, getHttpUrlConfigById, encodedUrl)
	var i HttpMonitConfig
	err := row.Scan(
		&i.EncodedUrl,
		&i.Url,
		&i.FriendlyName,
		&i.Interval,
		&i.Retries,
		&i.Timeout,
		&i.UpsideDown,
		&i.MaxRedirects,
		&i.Method,
		&i.AcceptedCodes,
		&i.BodyEncoding,
		&i.Body,
		&i.Headers,
		&i.AuthenticationMode,
		&i.ExpectedResponse,
	)
	return i, err
}

const getOneTimeTokenForServerRegistration = `-- name: GetOneTimeTokenForServerRegistration :one
SELECT one_time_token FROM servers WHERE id = ?
`

func (q *Queries) GetOneTimeTokenForServerRegistration(ctx context.Context, id string) (*string, error) {
	row := q.queryRow(ctx, q.getOneTimeTokenForServerRegistrationStmt, getOneTimeTokenForServerRegistration, id)
	var one_time_token *string
	err := row.Scan(&one_time_token)
	return one_time_token, err
}

const getServerByGidSid = `-- name: GetServerByGidSid :one
SELECT id, group_id, ip, mac, reg_status, one_time_token, one_time_token_expiry, name, "desc", fqdn, agent_version, os, arch, nats_subject, monit_enabled, notifs_enabled FROM servers WHERE group_id = ? AND id = ?
`

type GetServerByGidSidParams struct {
	GroupID int64  `json:"group_id"`
	ID      string `json:"id"`
}

func (q *Queries) GetServerByGidSid(ctx context.Context, arg GetServerByGidSidParams) (Server, error) {
	row := q.queryRow(ctx, q.getServerByGidSidStmt, getServerByGidSid, arg.GroupID, arg.ID)
	var i Server
	err := row.Scan(
		&i.ID,
		&i.GroupID,
		&i.Ip,
		&i.Mac,
		&i.RegStatus,
		&i.OneTimeToken,
		&i.OneTimeTokenExpiry,
		&i.Name,
		&i.Desc,
		&i.Fqdn,
		&i.AgentVersion,
		&i.Os,
		&i.Arch,
		&i.NatsSubject,
		&i.MonitEnabled,
		&i.NotifsEnabled,
	)
	return i, err
}

const getServerGroup = `-- name: GetServerGroup :one
SELECT id, name, desc FROM server_groups WHERE id = ?
`

func (q *Queries) GetServerGroup(ctx context.Context, id int64) (ServerGroup, error) {
	row := q.queryRow(ctx, q.getServerGroupStmt, getServerGroup, id)
	var i ServerGroup
	err := row.Scan(&i.ID, &i.Name, &i.Desc)
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

const listHttpUrlConfigs = `-- name: ListHttpUrlConfigs :many
SELECT encoded_url, url, friendly_name, interval, retries, timeout, upside_down, max_redirects, method, accepted_codes, body_encoding, body, headers, authentication_mode, expected_response FROM http_monit_configs LIMIT ? OFFSET ?
`

type ListHttpUrlConfigsParams struct {
	Limit  int64 `json:"limit"`
	Offset int64 `json:"offset"`
}

func (q *Queries) ListHttpUrlConfigs(ctx context.Context, arg ListHttpUrlConfigsParams) ([]HttpMonitConfig, error) {
	rows, err := q.query(ctx, q.listHttpUrlConfigsStmt, listHttpUrlConfigs, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []HttpMonitConfig{}
	for rows.Next() {
		var i HttpMonitConfig
		if err := rows.Scan(
			&i.EncodedUrl,
			&i.Url,
			&i.FriendlyName,
			&i.Interval,
			&i.Retries,
			&i.Timeout,
			&i.UpsideDown,
			&i.MaxRedirects,
			&i.Method,
			&i.AcceptedCodes,
			&i.BodyEncoding,
			&i.Body,
			&i.Headers,
			&i.AuthenticationMode,
			&i.ExpectedResponse,
		); err != nil {
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

const listServerGroups = `-- name: ListServerGroups :many
SELECT id, name, desc FROM server_groups
`

func (q *Queries) ListServerGroups(ctx context.Context) ([]ServerGroup, error) {
	rows, err := q.query(ctx, q.listServerGroupsStmt, listServerGroups)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ServerGroup{}
	for rows.Next() {
		var i ServerGroup
		if err := rows.Scan(&i.ID, &i.Name, &i.Desc); err != nil {
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

const listServersByGroup = `-- name: ListServersByGroup :many
SELECT id, group_id, ip, mac, reg_status, one_time_token, one_time_token_expiry, name, "desc", fqdn, agent_version, os, arch, nats_subject, monit_enabled, notifs_enabled FROM servers WHERE group_id = ?
`

func (q *Queries) ListServersByGroup(ctx context.Context, groupID int64) ([]Server, error) {
	rows, err := q.query(ctx, q.listServersByGroupStmt, listServersByGroup, groupID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Server{}
	for rows.Next() {
		var i Server
		if err := rows.Scan(
			&i.ID,
			&i.GroupID,
			&i.Ip,
			&i.Mac,
			&i.RegStatus,
			&i.OneTimeToken,
			&i.OneTimeTokenExpiry,
			&i.Name,
			&i.Desc,
			&i.Fqdn,
			&i.AgentVersion,
			&i.Os,
			&i.Arch,
			&i.NatsSubject,
			&i.MonitEnabled,
			&i.NotifsEnabled,
		); err != nil {
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

const updateOneTimeTokenForServerRegistration = `-- name: UpdateOneTimeTokenForServerRegistration :exec
UPDATE servers SET one_time_token = ?, reg_status = 'PENDING' WHERE id = ?
`

type UpdateOneTimeTokenForServerRegistrationParams struct {
	OneTimeToken *string `json:"one_time_token"`
	ID           string  `json:"id"`
}

func (q *Queries) UpdateOneTimeTokenForServerRegistration(ctx context.Context, arg UpdateOneTimeTokenForServerRegistrationParams) error {
	_, err := q.exec(ctx, q.updateOneTimeTokenForServerRegistrationStmt, updateOneTimeTokenForServerRegistration, arg.OneTimeToken, arg.ID)
	return err
}
