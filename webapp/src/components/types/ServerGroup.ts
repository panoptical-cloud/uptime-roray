export enum ServerRegStatusEnum {
  New = "NEW",
  Pending = "PENDING",
  Active = "ACTIVE",
}

export type Server = {
  id: string
  name: string
  fqdn: string
  ip: string
  agent_port: number
  agent_version: string
  reg_status: ServerRegStatusEnum
  desc: string
}

export type ServerGroup = {
  id: number
  name: string
  desc: string
}